package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	//Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//Connection to database
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		dbUsername, dbPassword, dbHost, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	return db, nil
}

func createTable(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		modified_at TIMESTAMPTZ DEFAULT NOW(),
		CONSTRAINT users_password_not_null CHECK (password <> ''),
		CONSTRAINT proper_email CHECK (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$')
		);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}
