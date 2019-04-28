package task

import (
	"context"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/robfig/cron"
)

func NewTask() object.ITask {
	return &task{}
}

type task struct {
	id       string
	taskType object.TaskType
	config   object.ITaskConfig
	cron     *cron.Cron
	running  bool
	err      error

	ctx    context.Context
	cancel func()
}

func (t *task) GetTaskId() string {
	return t.id
}

func (t *task) Start() {
	//TODO
}

func (t *task) Stop() {
	//TODO
}
