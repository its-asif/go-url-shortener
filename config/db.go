package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	dbURL := os.Getenv("DB_URL")
	var err error
	DB, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("DB Connection Error: %v", err)
	}
	fmt.Println("Connected to Database")
}
