package tasks

import "fmt"

type teamLead struct {
	Id   int
	Name string
}

func NewTeamLead(name string) *teamLead {
	return &teamLead{
		Id:   ID,
		Name: name,
	}
}

func (tl teamLead) DeligateTask(taskId task) (task, error) {

	for i, v := range TaskList {
		if v.Id == taskId.Id {
			if TaskList[i].Status == TestStatus {
				TaskList[i].Status = DoneStatus
			} else {
				TaskList[i].Status = DevStatus
			}
			return TaskList[i], nil
		}
	}

	return task{}, fmt.Errorf("task %d not exists", taskId.Id)
}
