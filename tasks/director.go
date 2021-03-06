package tasks

var ID int

type director struct {
	Id   int
	Name string
}

func NewDirector(name string) *director {
	return &director{
		Id:   ID,
		Name: name,
	}
}

func (dr director) GiveTask(taskid int, name, deadline string, tl *teamLead) (task, error) {
	newTask := task{
		Id:       taskid,
		Name:     name,
		Status:   InitStatus,
		Deadline: deadline,
		TeamLead: *tl,
	}

	TaskList = append(TaskList, newTask)
	return newTask, nil
}
