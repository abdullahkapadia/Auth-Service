package database

import (
	"context"
	"log"
)

func Migrate() error{

	query:= `
	CREATE TABLE IF NOT EXISTS users(
		id UUID PRIMARY KEY,

		name TEXT NOT NULL,

		email TEXT UNIQUE NOT NULL,

		password TEXT NOT NULL,

		created_at TIMESTAMP DEFAULT NOW(),

		updated_at TIMESTAMP DEFAULT NOW()

	);	
	`
	_, err := DB.Exec(context.Background(),query)

	if err != nil{
		return  err
	}
	log.Println("users table ready")

	return nil
}