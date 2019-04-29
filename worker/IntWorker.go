package worker

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV6/notify"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
	"strconv"
	"strings"
	"time"
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
	if w.config == nil {
		errMsg := fmt.Sprintf("int worker error: config is nil")
		log.Error(errMsg)
		w.chErr <- errors.New(errMsg)
		return
	}
	//get num
	num, err := w.getCheckNum()
	if err != nil {
		errMsg := fmt.Sprintf("get check num error: %s", err.Error())
		log.Error(errMsg)
		w.chErr <- errors.New(errMsg)
		return
	}

	//send msg
	var msg string
	if num >= w.config.FCheckMax || num <= w.config.FCheckMin {
		comm := common{}
		msg = comm.getMsg(w.config.FMsgTitle, strings.Replace(w.config.FMsgContent, "title", strconv.Itoa(num), -1))
		msg = strings.Trim(msg, " ")
		if msg != "" {
			msg = w.formatMsg(msg)
			//TODO send check
			_, _, _, _, _ = notify.SendMsg(w.config.FId, msg)
		}
	}

	//TODO save his data
	//TODO check action
}

func (w *intWorker) ClearHisData() {
	//TODO ClearHisData
	log.Debug("ClearHisData")
}

func (w *intWorker) GetChErr() <-chan error {
	return w.chErr
}

func (w *intWorker) formatMsg(msg string) string {
	if msg != "" {
		msg = goToolCommon.GetDateTimeStr(time.Now()) + "\n" + msg
	}
	return msg
}

//获取待检测值
func (w *intWorker) getCheckNum() (int, error) {
	rows, err := w.getRowsBySQL(w.config.FSearch)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rows.Close()
	}()
	list := make([]int, 0)
	var num int
	for rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			log.Error(err.Error())
			break
		} else {
			list = append(list, num)
		}
	}
	if err != nil {
		return 0, err
	}
	if len(list) != 1 {
		errMsg := fmt.Sprintf("SQL返回数量异常，exp:1,act:%d", len(list))
		log.Error(errMsg)
		return 0, errors.New(errMsg)
	}
	return list[0], nil
}

//查询数据
func (w *intWorker) getRowsBySQL(sql string) (*sql.Rows, error) {
	conn, err := goToolMSSql.GetConn(w.getDBConfig())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	rows, err := conn.Query(sql)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return rows, nil
}

//获取DB配置
func (w *intWorker) getDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: w.config.FServer,
		Port:   w.config.FPort,
		DbName: w.config.FDbName,
		User:   w.config.FDbUser,
		Pwd:    w.config.FDbPwd,
	}
}
