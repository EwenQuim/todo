package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/EwenQuim/todo-app/app/controllers"
	"github.com/EwenQuim/todo-app/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:generate yarn --cwd frontend build

//go:embed frontend/build/*
var reactBuild embed.FS

func main() {
	fsub, err := fs.Sub(reactBuild, "frontend/build")
	if err != nil {
		log.Fatal(err)
	}

	dab := database.InitDatabase("todo.db")

	s := database.Service{DB: dab, Regex: *regexp.MustCompile(`^ *([\w ]+) *: *(.*) *$`)}

	app := fiber.New()

	controllers.RegisterRoutes(app, s)

	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			return !strings.HasPrefix(c.Path(), "/static")
		},
		Level: compress.LevelBestSpeed,
	}))

	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(fsub),
		NotFoundFile: "index.html",
	}))
	app.Use(cors.New())
	// app.Static("", "./frontend/build") // For not embedding the build folder
	app.Use(logger.New())

	log.Fatal(app.Listen(":8083"))
}
