package database

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/EwenQuim/todo-app/app/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service struct {
	DB    *gorm.DB
	Regex regexp.Regexp
}

func InitDatabase(db_name string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	// Opens Database
	db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	if err := db.AutoMigrate(&model.Item{}, &model.Todo{}, &model.Tag{}); err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Database Migrated")
	return db
}
