package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kiki-ki/lesson-ent/ent"
)

func main() {
	db := NewDB()
	defer db.Conn.Close()

	// オートマイグレーションの実行
	if err := db.Conn.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

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
	Conn *ent.Client
}

func NewDB() *DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	entDB, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed openning connection to mysql: %v", err))
	}
	env := os.Getenv("ENV")

	// デバッグモードを利用
	if env != "staging" && env != "production" {
		entDB = entDB.Debug()
	}

	db := &DB{Conn: entDB}

	return db
}
