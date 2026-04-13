package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool
func DBUrl() string{
return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

}

func ConnectDB() {
	pool, err := pgxpool.New(context.Background(), DBUrl())
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("DB ping failed", err)
	}
	log.Println("Connected to PostgreSQL")
	DB=pool
	log.Println("Database connected successfully")
}


func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}