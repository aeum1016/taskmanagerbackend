package models

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var ( 
	Connection *pgxpool.Pool
	pgOnce sync.Once
)

func DBConnection() *pgxpool.Pool {
	pgOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file, ", err)
		}

		host := os.Getenv("DB_URL")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		dbname := os.Getenv("DB_NAME")

		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
			user, pass, host, port, dbname)
		
		Connection, err = pgxpool.New(context.Background(), connectionString)
		
		if err != nil {
			fmt.Printf("Failed to connect to database: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println("Connection established to database ", dbname)
		}
	})

	return Connection
}
