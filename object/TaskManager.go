package object

import (
	"fmt"
	"github.com/robfig/cron"
	"sync"
)

import log "github.com/Deansquirrel/goToolLog"

//taskManager对外接口对象
type ITaskManager interface {
}

type ITask interface {
	GetTaskId() string
	Start()
	Stop()
}

func NewTaskManager() ITaskManager {
	t := taskManager{
		taskList:     make(map[string]ITask),
		chRegister:   make(chan ITask),
		chUnregister: make(chan string),
	}
	t.start()
	return &t
}

//Task管理类
//作用：
// - 管理Task对象（缓存，添加，删除）
// - 定期刷新配置（根据配置，添加或删除任务）
// - 历史数据清理
type taskManager struct {
	taskList     map[string]ITask
	chRegister   chan ITask
	chUnregister chan string

	lock sync.Mutex
}

func (tm *taskManager) start() {

}

func (tm *taskManager) register(task ITask) {
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
	task, ok := tm.taskList[id]
	if !ok {
		log.Warn(fmt.Sprintf("task %s is not exist", id))
		return
	}
	task.Stop()
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
	}()
}

func (tm *taskManager) refreshConfig() {

}

func (tm *taskManager) clearHisData() {

}
