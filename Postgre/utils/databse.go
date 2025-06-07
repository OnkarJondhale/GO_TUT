package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDb() {
	_ = godotenv.Load(".env")

	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var err error
	DB, err = sql.Open("pgx", DATABASE_URL)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to reach database:", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}
