package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"regexp"

	"github.com/EwenQuim/todo-app/app/controllers"
	"github.com/EwenQuim/todo-app/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	chicors "github.com/go-chi/cors"
)

//go:generate yarn --cwd frontend build

//go:embed frontend/build/*
var reactBuild embed.FS

func main() {
	// Custom path to db
	var dbPath string
	flag.StringVar(&dbPath, "db", "todo.db", "path to database")
	flag.Parse()

	fsub, err := fs.Sub(reactBuild, "frontend/build")
	if err != nil {
		log.Fatal(err)
	}

	s := database.Service{
		DB:    database.InitDatabase(dbPath),
		Regex: *regexp.MustCompile(`^ *([\w ]+) *: *(.*) *$`),
	}
	r := chi.NewRouter()

	// Middleware stack
	r.Use(middleware.Compress(5, "text/html", "text/css"))

	r.Use(chicors.Handler(chicors.Options{
		AllowedOrigins: []string{"*"},
	}))
	r.Use(middleware.Logger)

	// Routes
	r.Get("/yo-{test}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yo" + chi.URLParam(r, "test")))
	})

	res := controllers.TodoResources{Service: s}
	res.RegisterRoutes(r)
	r.Handle("/*", http.FileServer(http.FS(fsub)))

	fmt.Println("server started at :8084")
	http.ListenAndServe(":8084", r)
}
