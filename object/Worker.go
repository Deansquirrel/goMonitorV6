package object

type IWorker interface {
	GetMsg()
	ClearHisData()
	GetChErr() <-chan error
}
