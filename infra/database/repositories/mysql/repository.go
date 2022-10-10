package mysql

import (
	database "api-user/infra/database"
	"database/sql"
	"log"
)

type Repositories struct{}

type RepositoriesInterface interface {
	CreateTransaction() *sql.Tx
}

func (*Repositories) CreateTransaction() (tx *sql.Tx) {
	conn := database.GetConnection()
	tx, err := conn.Begin()

	if err != nil {
		log.Panicln("TX Create: ", err)
	}

	return
}
