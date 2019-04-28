package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV6/global"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goToolMSSql"
)

import log "github.com/Deansquirrel/goToolLog"

func GetConfigList(taskType object.TaskType) []object.ITaskConfig {
	//TODO 获取配置列表
	list := make([]object.ITaskConfig, 0)
	return list
}

type Common struct {
}

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
