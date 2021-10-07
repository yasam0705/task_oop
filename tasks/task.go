package tasks

const (
	InitStatus = "initital"
	DevStatus  = "develop"
	TestStatus = "testing"
	DoneStatus = "done"
)

type task struct {
	Id       int
	Name     string
	Status   string
	Deadline string
	TeamLead teamLead
}

var TaskList []task = make([]task, 0)
