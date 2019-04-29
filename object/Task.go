package object

import (
	"database/sql"
	"fmt"
	"reflect"
)

import log "github.com/Deansquirrel/goToolLog"

//taskManager对外接口对象
type ITaskManager interface {
	GetTaskIdList() []string
	GetTask(id string) ITask
	GetTypeTaskIdList(taskType TaskType) []string
	GetChRegister() chan<- ITask
	GetChUnregister() chan<- string
}

type ITask interface {
	GetTaskId() string
	GetTaskType() TaskType
	GetTaskConfig() ITaskConfig
	Start()
	Stop()
	ClearHisData()
}

type ITaskConfig interface {
	GetTaskId() string
	GetSpec() string
	IsEqual(c ITaskConfig) bool
}

type IntTaskConfig struct {
	FId         string
	FServer     string
	FPort       int
	FDbName     string
	FDbUser     string
	FDbPwd      string
	FSearch     string
	FCron       string
	FCheckMax   int
	FCheckMin   int
	FMsgTitle   string
	FMsgContent string
}

func (config *IntTaskConfig) GetSpec() string {
	return config.FCron
}

func (config *IntTaskConfig) GetTaskId() string {
	return config.FId
}

func (config *IntTaskConfig) IsEqual(c ITaskConfig) bool {
	switch reflect.TypeOf(c).String() {
	case reflect.TypeOf(config).String():
		c, ok := c.(*IntTaskConfig)
		if !ok {
			return false
		}
		if config.FId != c.FId ||
			config.FServer != c.FServer ||
			config.FPort != c.FPort ||
			config.FDbName != c.FDbName ||
			config.FDbUser != c.FDbUser ||
			config.FDbPwd != c.FDbPwd ||
			config.FSearch != c.FSearch ||
			config.FCron != c.FCron ||
			config.FCheckMax != c.FCheckMax ||
			config.FCheckMin != c.FCheckMin ||
			config.FMsgTitle != c.FMsgTitle ||
			config.FMsgContent != c.FMsgContent {
			return false
		}
		return true
	default:
		return false
	}
}

type ITaskConfigRepositoryResource interface {
	GetSqlGetConfig() string
	GetSqlGetConfigList() string
	DataWrapper(rows *sql.Rows) ([]ITaskConfig, error)
}

//Sql IntTaskConfig
const (
	sqlGetIntTaskConfig = "" +
		"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
		"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
		"B.[FMsgTitle],B.[FMsgContent]" +
		" FROM [MConfig] A" +
		" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]"

	sqlGetIntTaskConfigById = "" +
		"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
		"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
		"B.[FMsgTitle],B.[FMsgContent]" +
		" FROM [MConfig] A" +
		" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]" +
		" WHERE B.[FId]=?"
)

type IntTaskConfigRepositoryResource struct{}

func (IntTaskConfigRepositoryResource) GetSqlGetConfig() string {
	return sqlGetIntTaskConfigById
}

func (IntTaskConfigRepositoryResource) GetSqlGetConfigList() string {
	return sqlGetIntTaskConfig
}

func (IntTaskConfigRepositoryResource) DataWrapper(rows *sql.Rows) ([]ITaskConfig, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fServer, fDbName, fDbUser, fDbPwd, fSearch, fCron, fMsgTitle, fMsgContent string
	var fPort, fCheckMax, fCheckMin int
	resultList := make([]ITaskConfig, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(
			&fId, &fServer, &fPort, &fDbName, &fDbUser,
			&fDbPwd, &fSearch, &fCron, &fCheckMax, &fCheckMin,
			&fMsgTitle, &fMsgContent)
		if err != nil {
			log.Error(fmt.Sprintf("convert data error: %s", err.Error()))
			return nil, err
		}
		config := IntTaskConfig{
			FId:         fId,
			FServer:     fServer,
			FPort:       fPort,
			FDbName:     fDbName,
			FDbUser:     fDbUser,
			FDbPwd:      fDbPwd,
			FSearch:     fSearch,
			FCron:       fCron,
			FCheckMax:   fCheckMax,
			FCheckMin:   fCheckMin,
			FMsgTitle:   fMsgTitle,
			FMsgContent: fMsgContent,
		}
		resultList = append(resultList, &config)
	}
	if rows.Err() != nil {
		log.Error(fmt.Sprintf("convert data error: %s", rows.Err().Error()))
		return nil, rows.Err()
	}
	return resultList, nil
}
