package main

import (
	"fmt"
	"oop_task/tasks"
)

func main() {
	dr1 := tasks.NewDirector("Director name")
	tl1 := tasks.NewTeamLead("TeamLead name")
	d := tasks.NewDeveloper("Developer name")

	newTask, _ := dr1.GiveTask(10, "create corparative website", "11.10.2021")
	fmt.Println(newTask)

	newTask, err := tl1.DeligateTask(newTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(newTask)

	newTask, err = d.DevelopTask(newTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(newTask)

	newTask, err = tl1.DeligateTask(newTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(newTask)
}
