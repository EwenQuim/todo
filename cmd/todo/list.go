package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func list() func(c *cli.Context) error {
	return func(c *cli.Context) error {

		todo := sortSpecial(getTodo(URL))

		var list []string
		var selectedBeforeChange []string

		for _, t := range todo.Items {
			list = append(list, t.Content)
			if t.Done {
				selectedBeforeChange = append(selectedBeforeChange, t.Content)
			}
		}

		p := &survey.MultiSelect{
			Message: "TODO: " + todo.Title,
			Options: list,
			Default: selectedBeforeChange,
		}

		// answers := struct {
		// 	List []string `survey:"list"` // or you can tag fields to match a specific name
		// }{}
		answers := []string{}

		err := survey.AskOne(p, &answers)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		var wg sync.WaitGroup
		for _, t := range todo.Items {
			t.toMarkAsDone = contains(answers, t.Content)
			if t.toMarkAsDone != t.Done {
				wg.Add(1)
				go func(t localItem) {
					defer wg.Done()
					_, err := http.Get(URL + "/" + strconv.Itoa(int(t.ID)) + "/switch")
					if err != nil {
						fmt.Println(err)
					}
				}(t)
			}
		}
		wg.Wait()
		return nil
	}
}
