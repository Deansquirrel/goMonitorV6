package task

import (
	"context"
	"fmt"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goMonitorV6/repository"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/robfig/cron"
	"sync"
)

import log "github.com/Deansquirrel/goToolLog"

func NewTaskManager() object.ITaskManager {
	t := taskManager{
		taskList:     make(map[string]object.ITask),
		chRegister:   make(chan object.ITask),
		chUnregister: make(chan string),
	}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	t.start()
	return &t
}

//Task管理类
//作用：
// - 管理Task对象（缓存，添加，删除）
// - 定期刷新配置（根据配置，添加或删除任务）
// - 历史数据清理
type taskManager struct {
	taskList     map[string]object.ITask
	chRegister   chan object.ITask
	chUnregister chan string

	ctx    context.Context
	cancel func()

	lock sync.Mutex
}

func (tm *taskManager) GetTaskIdList() []string {
	list := make([]string, 0)
	for id := range tm.taskList {
		list = append(list, id)
	}
	return list
}

func (tm *taskManager) GetTask(id string) object.ITask {
	t, ok := tm.taskList[id]
	if ok {
		return t
	}
	return nil
}

func (tm *taskManager) GetTypeTaskIdList(taskType object.TaskType) []string {
	list := make([]string, 0)
	for id, t := range tm.taskList {
		if t.GetTaskType() == taskType {
			list = append(list, id)
		}
	}
	return list
}

func (tm *taskManager) GetChRegister() chan<- object.ITask {
	return tm.chRegister
}

func (tm *taskManager) GetChUnregister() chan<- string {
	return tm.chUnregister
}

func (tm *taskManager) start() {
	tm.startRegularJob()
	go func() {
		for {
			select {
			case t := <-tm.chRegister:
				tm.register(t)
			case id := <-tm.chUnregister:
				tm.unregister(id)
			case <-tm.ctx.Done():
				return
			}
		}
	}()
}

func (tm *taskManager) register(task object.ITask) {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	_, ok := tm.taskList[task.GetTaskId()]
	if ok {
		log.Warn(fmt.Sprintf("task %s is already exist", task.GetTaskId()))
		return
	}
	tm.taskList[task.GetTaskId()] = task
	task.Start()
	log.Info(fmt.Sprintf("task %s is added", task.GetTaskId()))
	return
}

func (tm *taskManager) unregister(id string) {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	t, ok := tm.taskList[id]
	if !ok {
		log.Warn(fmt.Sprintf("task %s is not exist", id))
		return
	}
	t.Stop()
	delete(tm.taskList, id)
}

func (tm *taskManager) startRegularJob() {
	go func() {
		c := cron.New()
		var err error
		err = c.AddFunc("0 0 * * * ?", tm.refreshConfig)
		if err != nil {
			log.Error(fmt.Sprintf("add refresh config task error [%s]", err.Error()))
		}
		err = c.AddFunc("0 0 0 * * ?", tm.clearHisData)
		if err != nil {
			log.Error(fmt.Sprintf("add clear his data task error [%s]", err.Error()))
		}
		c.Start()
		select {
		case <-tm.ctx.Done():
			c.Stop()
			return
		}
	}()
}

func (tm *taskManager) refreshConfig() {
	for _, taskType := range object.TaskTypeList {
		repConfigList := repository.GetConfigList(taskType)
		repConfigMap := make(map[string]object.ITaskConfig, 0)
		repConfigIdList := make([]string, 0)
		for _, repConfig := range repConfigList {
			repConfigIdList = append(repConfigIdList, repConfig.GetTaskId())
			repConfigMap[repConfig.GetTaskId()] = repConfig
		}
		cacheTaskIdList := tm.GetTypeTaskIdList(taskType)
		addList, delList, checkList := goToolCommon.CheckDiff(repConfigIdList, cacheTaskIdList)
		for _, id := range addList {
			newConfig, ok := repConfigMap[id]
			if ok {
				tm.GetChRegister() <- NewTask(taskType, newConfig)
			}
		}
		for _, id := range delList {
			tm.GetChUnregister() <- id
		}
		for _, id := range checkList {
			newConfig, ok := repConfigMap[id]
			if ok {
				oldTask := tm.GetTask(id)
				if !oldTask.GetTaskConfig().IsEqual(newConfig) {
					tm.GetChUnregister() <- id
					tm.GetChRegister() <- NewTask(taskType, newConfig)
				}
			}
		}
	}
}

func (tm *taskManager) clearHisData() {
	for _, t := range tm.taskList {
		t.ClearHisData()
	}
}
