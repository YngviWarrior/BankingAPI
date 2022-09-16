package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct{}

type DatabaseInterface interface {
	CreateMysqlPool() *sql.DB
	CreatePostgresPool() *sql.DB
}

func GetConnection() *sql.DB {
	var db DatabaseInterface = &Database{}
	return db.CreateMysqlPool()
}

func (d *Database) CreateMysqlPool() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB"))

	if err != nil {
		log.Fatal("DC 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func (d *Database) CreatePostgresPool() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB"))

	if err != nil {
		log.Fatal("DC 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
