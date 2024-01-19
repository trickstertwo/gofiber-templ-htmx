package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getConnection() {
	var err error

	if db != nil {
		return
	}

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)

	// Init PostgreSQL database
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("🔥 failed to ping the database: %s", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")
}

func MakeMigrations() {
	getConnection()

	stmt := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			username VARCHAR(64) NOT NULL
		)
	`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	stmt = `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			created_by INTEGER NOT NULL,
			title VARCHAR(64) NOT NULL,
			description VARCHAR(255),
			status BOOLEAN DEFAULT FALSE,
			FOREIGN KEY (created_by) REFERENCES users(id)
		)
	`

	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}
