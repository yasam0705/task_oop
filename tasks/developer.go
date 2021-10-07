package tasks

import (
	"fmt"
)

type developer struct {
	Id       int
	Name     string
	TeamLead teamLead
}

func NewDeveloper(name string, t *teamLead) *developer {
	return &developer{
		Id:       ID,
		Name:     name,
		TeamLead: *t,
	}
}

func (d developer) DevelopTask(taskId task) (task, error) {

	for i, v := range TaskList {
		if v.Id == taskId.Id {
			TaskList[i].Status = TestStatus
			return TaskList[i], nil
		}
	}

	return task{}, fmt.Errorf("task %d not exists", taskId.Id)
}
