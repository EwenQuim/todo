package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

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

	// Custom portNumber
	var portNumber int
	flag.IntVar(&portNumber, "port", 8084, "port to listen on")
	flag.Parse()

	port := fmt.Sprintf(":%d", portNumber)

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
	r.Use(middleware.Compress(5, "text/html", "text/javascript", "text/css", "application/javascript"))

	r.Use(chicors.Handler(chicors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8084"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	r.Use(middleware.Logger)

	// Cache
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				if strings.HasSuffix(r.URL.Path, ".js") || strings.HasSuffix(r.URL.Path, ".css") {
					w.Header().Set("Cache-Control", "public, max-age=86400, stale-while-revalidate=604800")
				}
			}
			h.ServeHTTP(w, r)
		})
	})

	// Routes
	res := controllers.TodoResources{Service: s}
	res.RegisterRoutes(r)

	r.Handle("/*", http.FileServer(spaFileSystem{http.FS(fsub)}))

	fmt.Printf("server started on http://localhost%s\n", port)
	http.ListenAndServe(port, r)
}

type spaFileSystem struct {
	root http.FileSystem
}

func (fs spaFileSystem) Open(name string) (http.File, error) {
	f, err := fs.root.Open(name)
	if os.IsNotExist(err) {
		return fs.root.Open("index.html")
	}
	return f, err
}
