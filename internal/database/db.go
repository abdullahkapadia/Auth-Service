package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/adullahkapadia/auth-service/internal/config"
)

var DB *pgxpool.Pool

func Connect(cfg *config.Config) error {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return err
	}

	DB = db

	log.Println("Database Connected Successfully")

	return nil
}