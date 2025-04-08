package db

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
	"microservices/ent"
	"microservices/ent/migrate"
	"microservices/pkg/config"
)

type Client struct {
	*ent.Client
}

func New(config *config.Config) (*Client, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.DB.Host, config.DB.Port, config.DB.User,
		config.DB.Name, config.DB.Password, config.DB.SSLMode)
	fmt.Println(dsn)
	db, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		return nil, err
	}
	client := &Client{db}
	if err = migrateDB(client); err != nil {
		return nil, err
	}
	return client, nil
}

func migrateDB(client *Client) error {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false),
	); err != nil {
		return err
	}
	return nil
}
