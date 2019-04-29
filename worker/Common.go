package worker

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV6/object"
	"reflect"
	"strings"
)

import log "github.com/Deansquirrel/goToolLog"

func GetWorker(config object.ITaskConfig) object.IWorker {
	switch reflect.TypeOf(config).String() {
	case reflect.TypeOf(&object.IntTaskConfig{}).String():
		return NewIntWorker(config.(*object.IntTaskConfig))
	default:
		log.Error(fmt.Sprintf("unexpected task config type,got %s", reflect.TypeOf(config).String()))
		return nil
	}
}

type common struct{}

//获取待发送消息
func (c *common) getMsg(title, content string) string {
	msg := ""
	titleList := strings.Split(title, "###")
	if len(titleList) > 0 {
		for _, t := range titleList {
			if strings.Trim(t, " ") != "" {
				if msg != "" {
					msg = msg + "\n"
				}
				msg = msg + strings.Trim(t, " ")
			}
		}
	}
	contentList := strings.Split(content, "###")
	if len(contentList) > 0 {
		for _, t := range contentList {
			if strings.Trim(t, " ") != "" {
				if msg != "" {
					msg = msg + "\n"
				}
				msg = msg + strings.Trim(t, " ")
			}
		}
	}
	return msg
}
