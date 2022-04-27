package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/urfave/cli/v2"
)

func clean() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Println("Cleaning up todo list")
		todo := getTodo(URL)

		var wg sync.WaitGroup
		for _, task := range todo.Items {
			if task.Done {
				wg.Add(1)
				go func(task localItem) {
					defer wg.Done()

					_, err := http.Get(URL + "/delete/" + string(task.ID))
					if err != nil {
						fmt.Println(err)
					}
				}(task)
			}
		}
		wg.Wait()
		fmt.Println("Done")

		return nil
	}
}
