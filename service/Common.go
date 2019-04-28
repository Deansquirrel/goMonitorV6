package service

import (
	"github.com/Deansquirrel/goMonitorV6/global"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goToolCommon"
	"os"
	"os/signal"
	"syscall"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

var taskManager object.ITaskManager

func init() {
	taskManager = object.NewTaskManager()
}

//启动服务内容
func Start() error {

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
			defer global.Cancel()
		case <-global.Ctx.Done():
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			log.Debug(goToolCommon.Guid())
			time.Sleep(time.Second)
		}
		global.Cancel()
	}()
	return nil
}
