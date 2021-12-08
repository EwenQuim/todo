package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getTodo(url string) localTodo {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Make http GET request and print response to stdout
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var todo localTodo

	// Unmarchal resp.Body into todo
	if err := json.Unmarshal(body, &todo); err != nil {
		panic(err)
	}
	return todo
}
