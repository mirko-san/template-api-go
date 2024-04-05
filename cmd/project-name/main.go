package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}

type Config struct {
	db DBConfig
}

type DBConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func main() {
	dbConfig := DBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   "main",
	}
	config := Config{
		db: dbConfig,
	}

	dbDriver := "postgres"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.db.host,
		config.db.port,
		config.db.user,
		config.db.password,
		config.db.dbname,
	)

	db, err := setupDB(dbDriver, dsn)

	fmt.Println("Hello, world!")

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
