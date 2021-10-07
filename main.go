package main

import (
	"fmt"
	"oop_task/tasks"
)

func main() {
	dr1 := tasks.NewDirector("Director name")
	tl1 := tasks.NewTeamLead("TeamLead name", dr1)
	d := tasks.NewDeveloper("Developer name", tl1)

	newTask, _ := dr1.GiveTask(10, "create corporate website", "11.10.2021", tl1)
	fmt.Println(newTask)

	newTask, err := newTask.TeamLead.DeligateTask(newTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(newTask)

	newTask, err = d.DevelopTask(newTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(newTask)

	newTask, err = newTask.TeamLead.DeligateTask(newTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(newTask)
}
