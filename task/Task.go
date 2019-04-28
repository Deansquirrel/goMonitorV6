package task

import (
	"context"
	"fmt"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goMonitorV6/worker"
	"github.com/kataras/iris/core/errors"
	"github.com/robfig/cron"
)

import log "github.com/Deansquirrel/goToolLog"

func NewTask(taskType object.TaskType, config object.ITaskConfig) object.ITask {
	t := task{
		taskType: taskType,
		config:   config,
		cron:     nil,
		running:  false,
		err:      nil,
	}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	return &t
}

type task struct {
	taskType object.TaskType
	config   object.ITaskConfig
	cron     *cron.Cron
	running  bool
	err      error

	ctx    context.Context
	cancel func()
}

func (t *task) ClearHisData() {
	//TODO ClearHisData
	log.Debug("ClearHisData")
}

func (t *task) GetTaskConfig() object.ITaskConfig {
	return t.config
}

func (t *task) GetTaskType() object.TaskType {
	return t.taskType
}

func (t *task) GetTaskId() string {
	return t.config.GetTaskId()
}

func (t *task) Start() {
	if !t.running {
		t.running = true
		if t.config == nil {
			t.err = errors.New(fmt.Sprintf("start errpr: config is nil"))
			return
		}
		w := worker.GetWorker(t.config)
		if w == nil {
			t.err = errors.New(fmt.Sprintf("get worker error: worker is nil"))
			return
		}
		c := cron.New()
		err := c.AddFunc(t.config.GetSpec(), w.GetMsg)
		if err != nil {
			t.err = errors.New(fmt.Sprintf("add func error: %s", err.Error()))
			return
		}
		c.Start()
		go func() {
			select {
			case err := <-w.GetChErr():
				t.err = errors.New(fmt.Sprintf("task run error: %s", err.Error()))
			case <-t.ctx.Done():
				return
			}
		}()
	}
}

func (t *task) Stop() {
	if t.running {
		t.running = false
		t.cron.Stop()
		t.cancel()
	}
}
