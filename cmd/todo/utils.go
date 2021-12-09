package main

import (
	"sort"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sortSpecial(todo localTodo) localTodo {
	sort.Slice(todo.Items, func(i, j int) bool {
		a := todo.Items[i].Content
		b := todo.Items[j].Content
		if strings.Contains(a, ": ") && strings.Contains(b, ": ") {
			return strings.Compare(a, b) < 0
		}

		return !strings.Contains(a, ": ")
	})
	return todo
}
