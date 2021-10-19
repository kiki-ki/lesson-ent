package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kiki-ki/lesson-ent/ent"
	"github.com/kiki-ki/lesson-ent/ent/migrate"
)

type EntClient struct {
	*ent.Client
}

func NewEntClient() *EntClient {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		panic(fmt.Sprintf("failed openning connection to mysql: %v", err))
	}
	env := os.Getenv("ENV")

	// デバッグモードを利用
	if env != "staging" && env != "production" {
		client = client.Debug()
	}

	return &EntClient{client}
}

func (c *EntClient) Migrate() {
	err := c.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
