package service

import (
	"github.com/Deansquirrel/goMonitorV6/global"
	"github.com/Deansquirrel/goToolCommon"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

//启动服务内容
func Start() error {
	go func() {
		for i := 0; i < 5; i++ {
			log.Debug(goToolCommon.Guid())
			time.Sleep(time.Second)
		}
		global.Cancel()
	}()
	return nil
}
