package worker

import (
	"github.com/Deansquirrel/goMonitorV6/object"
	log "github.com/Deansquirrel/goToolLog"
)

func NewIntWorker(config *object.IntTaskConfig) *intWorker {
	return &intWorker{
		config: config,
		chErr:  make(chan error),
	}
}

type intWorker struct {
	config *object.IntTaskConfig
	chErr  chan error
}

func (w *intWorker) GetMsg() {
	//TODO GetMsg send msg check action
	log.Debug("GetMsg")
}

func (w *intWorker) ClearHisData() {
	//TODO ClearHisData
	log.Debug("ClearHisData")
}

func (w *intWorker) GetChErr() <-chan error {
	return w.chErr
}
