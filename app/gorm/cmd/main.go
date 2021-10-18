package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := NewDB()
	r := chi.NewRouter()
	setRouter(r, db)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func setRouter(r chi.Router, db *DB) {
	// logging
	r.Use(middleware.Logger)

	// routing
	// userController := controller.NewUserHandler(db)
	// r.Route("/users", func(r chi.Router) {
	// 	r.Get("/", userController.Index)
	// 	r.Get("/{userId}", userController.Show)
	// 	r.Put("/{userId}", userController.Update)
	// 	r.Post("/", userController.Create)
	// 	r.Delete("/{userId}", userController.Delete)
	// 	r.Get("/transaction", userController.Transaction)
	// })
}

type DB struct {
	Conn *gorm.DB
}

func NewDB() *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("GORM_DB_NAME"))

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db := &DB{Conn: gormDB}

	return db
}
