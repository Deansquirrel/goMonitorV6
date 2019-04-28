package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV6/object"
)

import log "github.com/Deansquirrel/goToolLog"

type IConfigRepository interface {
	GetConfig(id string) (object.ITaskConfig, error)
	GetConfigList() ([]object.ITaskConfig, error)
}

func NewConfigRepository(taskType object.TaskType) (IConfigRepository, error) {
	switch taskType {
	case object.Int:
		return &configRepository{&object.IntTaskConfigRepositoryResource{}}, nil
	case object.CrmDzXfTest:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	case object.Health:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	case object.WebState:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	default:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	}
}

type configRepository struct {
	resource object.ITaskConfigRepositoryResource
}

func (r *configRepository) GetConfig(id string) (object.ITaskConfig, error) {
	c := Common{}
	rows, err := c.GetRowsBySQL(r.resource.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(fmt.Sprintf("get config error,sql [%s],id [%s]", r.resource.GetSqlGetConfig(), id))
		return nil, err
	}
	list, err := r.resource.DataWrapper(rows)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	} else {
		return nil, nil
	}
}

func (r *configRepository) GetConfigList() ([]object.ITaskConfig, error) {
	c := Common{}
	rows, err := c.GetRowsBySQL(r.resource.GetSqlGetConfigList())
	if err != nil {
		log.Error(fmt.Sprintf("get config list error,sql [%s]", r.resource.GetSqlGetConfig()))
		return nil, err
	}
	list, err := r.resource.DataWrapper(rows)
	if err != nil {
		return nil, err
	}
	return list, nil
}
