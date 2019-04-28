package service

import (
	"github.com/Deansquirrel/goMonitorV6/global"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goMonitorV6/repository"
	"github.com/Deansquirrel/goMonitorV6/task"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var taskManager object.ITaskManager

func init() {
	taskManager = task.NewTaskManager()
}

//启动服务内容
func Start() error {

	//接收系统退出消息
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			os.Interrupt,
			syscall.SIGINT,
			os.Kill,
			syscall.SIGKILL,
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			time.Sleep(time.Second * 3)
			defer global.Cancel()
		case <-global.Ctx.Done():
		}
	}()

	//go func() {
	//	for i := 0; i < 5; i++ {
	//		log.Debug(goToolCommon.Guid())
	//		time.Sleep(time.Second)
	//	}
	//	global.Cancel()
	//}()

	for _, taskType := range object.TaskTypeList {
		startTask(taskType)
	}
	return nil
}

func startTask(taskType object.TaskType) {
	configList := repository.GetConfigList(taskType)
	for _, config := range configList {
		t := task.NewTask(taskType, config)
		taskManager.GetChRegister() <- t
	}
}
