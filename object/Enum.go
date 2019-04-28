package object

type TaskType int

const (
	Int TaskType = iota
	CrmDzXfTest
	Health
	WebState
)

var TaskTypeList []TaskType

func init() {
	TaskTypeList = make([]TaskType, 0)
	TaskTypeList = append(TaskTypeList, Int)
	TaskTypeList = append(TaskTypeList, CrmDzXfTest)
	TaskTypeList = append(TaskTypeList, Health)
	TaskTypeList = append(TaskTypeList, WebState)
}
