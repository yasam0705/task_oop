package tasks

import (
	"fmt"
)

type developer struct {
	Id   int
	Name string
}

func NewDeveloper(name string) *developer {
	return &developer{
		Id:   ID,
		Name: name,
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
