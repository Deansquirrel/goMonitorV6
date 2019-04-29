package repository

import (
	"database/sql"
	"fmt"
	"github.com/Deansquirrel/goMonitorV6/global"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goToolMSSql"
)

import log "github.com/Deansquirrel/goToolLog"

func GetTaskConfigList(taskType object.TaskType) []object.ITaskConfig {
	rep, err := NewConfigRepository(taskType)
	if err != nil {
		log.Error(fmt.Sprintf("get task config rep error: %s", err.Error()))
		return nil
	}
	if rep == nil {
		log.Error(fmt.Sprintf("get task config rep error: rep is nil"))
		return nil
	}
	list, err := rep.GetConfigList()
	if err != nil {
		log.Error(fmt.Sprintf("get task config list error: %s", err.Error()))
		return nil
	}
	return list
}

func GetNotifyConfig(notifyType object.NotifyType, id string) object.INotifyConfig {
	rep, err := NewNotifyRepository(notifyType)
	if err != nil {
		log.Error(fmt.Sprintf("get notify config rep error: %s", err.Error()))
		return nil
	}
	if rep == nil {
		log.Error(fmt.Sprintf("get notify config rep error: rep is nil"))
		return nil
	}
	config, err := rep.GetNotify(id)
	if err != nil {
		log.Error(fmt.Sprintf("get notify config list error: %s", err.Error()))
		return nil
	}
	return config
}

type Common struct{}

//获取配置库连接配置
func (r *Common) getConfigDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: global.SysConfig.DB.Server,
		Port:   global.SysConfig.DB.Port,
		DbName: global.SysConfig.DB.DbName,
		User:   global.SysConfig.DB.User,
		Pwd:    global.SysConfig.DB.Pwd,
	}
}

func (r *Common) GetRowsBySQL(sql string, args ...interface{}) (*sql.Rows, error) {
	conn, err := goToolMSSql.GetConn(r.getConfigDBConfig())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if args == nil {
		rows, err := conn.Query(sql)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		return rows, nil
	} else {
		rows, err := conn.Query(sql, args...)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		return rows, nil
	}
}

func (r *Common) SetRowsBySQL(sql string, args ...interface{}) error {
	conn, err := goToolMSSql.GetConn(r.getConfigDBConfig())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if args == nil {
		_, err = conn.Exec(sql)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		return nil
	} else {
		_, err := conn.Exec(sql, args...)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		return nil
	}
}
