package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var taskCommand = cli.Command{
	Name:   "task",
	Usage:  "list tasks uncomleted",
	Action: doTask,
}

func doTask(c *cli.Context) {

	w := newWLClient()

	lists, _ := w.List.GetAll()

	fmt.Println("== ALL Tasks Not Completed ==\n")
	for _, list := range lists {
		tasks, _ := w.Task.GetByListID(list.ID, false)
		for _, task := range tasks {
			fmt.Println("→ " + list.Title)
			star := "　"
			if task.Starred {
				star = "☆"
			}

			fmt.Printf("%10d %s % s %s\n", task.ID, star, task.Title, task.DueDate)
			fmt.Println()
		}
	}
}
