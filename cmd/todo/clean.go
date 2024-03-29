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
		todo, err := getTodo(URL)
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		for _, task := range todo.Items {
			if task.Done {
				wg.Add(1)
				go func(task localItem) {
					defer wg.Done()

					// Create client
					client := &http.Client{}

					// Create request
					req, err := http.NewRequest("DELETE", "http://www.example.com/bucket/sample", nil)
					if err != nil {
						fmt.Println(err)
						return
					}

					// Fetch Request
					resp, err := client.Do(req)
					if err != nil {
						fmt.Println(err)
						return
					}
					defer resp.Body.Close()
				}(task)
			}
		}
		wg.Wait()
		fmt.Println("Done")

		return nil
	}
}
