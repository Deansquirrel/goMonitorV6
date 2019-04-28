package worker

import "github.com/Deansquirrel/goMonitorV6/object"

type IWorker interface {
	GetMsg()
	ClearHisData()
	GetChErr() <-chan error
}

func GetWorker(config object.ITaskConfig) IWorker {
	//TODO GetWorker
	return nil
}
