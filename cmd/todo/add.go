package main

import (
	"errors"
	"net/http"

	"github.com/urfave/cli/v2"
)

func add() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if !c.Args().Present() {
			return errors.New("cannot add an empty task")
		}
		s := c.Args().Slice()
		content := ""
		for _, v := range s {
			content += v + " "
		}

		_, err := http.Get(URL + "/new?content=" + content)
		if err != nil {
			return errors.New("not connected to the server")
		}
		return nil
	}
}
