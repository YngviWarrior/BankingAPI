package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sql.DB
}

type DatabaseInterface interface {
	DbConnect() *Database
}

func (d *Database) DbConnect() *Database {
	db, err := sql.Open("mysql", os.Getenv("DB"))

	if err != nil {
		log.Fatal("DC 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	d.Conn = db

	return d
}
