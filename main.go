package main

import (
	"book/api"
	"book/storage"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	PostgresUser     = "postgres"
	PostgresPassword = "12345"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "book_db"
)

func main() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	storage := storage.NewDBManager(db)

	server := api.NewServer(storage)

	err = server.Run(":8003")
	if err != nil {
		log.Fatalf("failed to start sever: %v", err)
	}

}
