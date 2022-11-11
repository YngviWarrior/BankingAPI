package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Pool *sql.DB
}

type DatabaseInterface interface {
	CreatePool()
	CreateConnection() (tx *sql.Tx, conn *sql.Conn)
}

func (d *Database) CreatePool() {
	d.Pool = CreateMysqlPool()
}

func (d *Database) CreateConnection() (tx *sql.Tx, conn *sql.Conn) {
	ctx := context.TODO()

	conn, err := d.Pool.Conn(ctx)
	if err != nil {
		log.Panicln("Conn Create: ", err)
	}

	tx, err = conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Panicln("TX Create: ", err)
	}

	return
}

func CreateMysqlPool() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB"))

	if err != nil {
		log.Fatal("DC 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Second * 30)
	db.SetMaxIdleConns(500)

	return db
}

func CreatePostgresPool() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB"))

	if err != nil {
		log.Fatal("DC 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
