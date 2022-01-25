package main

import (
	"log"
	"os"

	"github.com/EwenQuim/todo-app/app/model"
	"github.com/urfave/cli/v2"
)

type localTodo struct {
	model.Todo
	Items []localItem
}

type localItem struct {
	model.Item
	toMarkAsDone bool
}

const URL = "https://todo.quimerch.com/api/todo/435ebc21-d12a-4979-a988-48143e31fca6"

func main() {
	app := &cli.App{
		Name:        "todo",
		Usage:       "A simple CLI program to manage your tasks",
		Description: "Can work locally or remotely. By default, it will work locally. To work remotely, you need to specify the URL of the API server.",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a", "new", "n"},
				Usage:   "add a task to the list",
				Action:  add(),
			},
			{
				Name:    "list",
				Aliases: []string{"ls", "l"},
				Usage:   "list the tasks and their status",
				Action:  list(),
			},
			{
				Name:    "clean",
				Aliases: []string{"clean", "c"},
				Usage:   "remove the tasks that are marked as done",
				Action:  clean(),
			},
		},
	}
	app.Action = list()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
