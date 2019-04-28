package object

import "reflect"

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
