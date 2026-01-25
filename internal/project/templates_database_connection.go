package project

func databaseConnectionTemplate() string {
	return `package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	// Carrega o .env, ignora erro se não existir
	_ = godotenv.Load()

	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	porta := os.Getenv("DB_PORT")
	banco := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		usuario,
		senha,
		host,
		porta,
		banco,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}
`
}
