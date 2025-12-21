package project

func databaseConnectionTemplate() string {
	return `package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	usuario := "root"
	senha := "senha"
	host := "127.0.0.1"
	porta := "3306"
	banco := "meubanco"

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
