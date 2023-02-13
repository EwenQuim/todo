package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getTodo(url string) (localTodo, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Make http GET request and print response to stdout
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var todo localTodo

	// Unmarchal resp.Body into todo
	if err := json.Unmarshal(body, &todo); err != nil {
		return localTodo{}, fmt.Errorf("error unmarshaling %s into todo: %v", string(body), err)
	}
	return todo, nil
}
