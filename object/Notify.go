package object

import (
	"database/sql"
	"fmt"
)

import log "github.com/Deansquirrel/goToolLog"

type INotify interface {
	SendMsg(msg string)
}

type INotifyConfig interface {
	GetNotifyId() string
}

type DingTalkRobotNotifyConfig struct {
	FId         string
	FWebHookKey string
	FAtMobiles  string
	FIsAtAll    int
}

func (notify *DingTalkRobotNotifyConfig) GetNotifyId() string {
	return notify.FId
}

type INotifyRepositoryResource interface {
	GetSqlGetNotify() string
	GetSqlGetNotifyList() string
	DataWrapper(rows *sql.Rows) ([]INotifyConfig, error)
}

//Sql DingTalkRobotNotify
const (
	sqlGetDingTalkRobot = "" +
		"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
		" FROM [NConfig] A" +
		" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]"

	sqlGetDingTalkRobotById = "" +
		"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
		" FROM [NConfig] A" +
		" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]" +
		" WHERE A.[FId]=?"
)

type DingTalkRobotNotifyRepositoryResource struct{}

func (DingTalkRobotNotifyRepositoryResource) GetSqlGetNotify() string {
	return sqlGetDingTalkRobotById
}

func (DingTalkRobotNotifyRepositoryResource) GetSqlGetNotifyList() string {
	return sqlGetDingTalkRobot
}

func (DingTalkRobotNotifyRepositoryResource) DataWrapper(rows *sql.Rows) ([]INotifyConfig, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fWebHookKey, fAtMobiles string
	var fIsAtAll int
	rList := make([]INotifyConfig, 0)
	for rows.Next() {
		err := rows.Scan(&fId, &fWebHookKey, &fAtMobiles, &fIsAtAll)
		if err != nil {
			log.Error(fmt.Sprintf("read rows data error: %s", err.Error()))
			return nil, err
		}
		config := DingTalkRobotNotifyConfig{
			FId:         fId,
			FWebHookKey: fWebHookKey,
			FAtMobiles:  fAtMobiles,
			FIsAtAll:    fIsAtAll,
		}
		rList = append(rList, &config)
	}
	if rows.Err() != nil {
		log.Error(fmt.Sprintf("read rows data error: %s", rows.Err().Error()))
		return nil, rows.Err()
	}
	return rList, nil
}
