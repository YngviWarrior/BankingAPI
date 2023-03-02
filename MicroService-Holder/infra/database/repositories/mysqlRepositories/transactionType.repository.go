package mysql

import (
	"database/sql"
	transactionType "holder-ms/core/entities/transactionType"
	"holder-ms/infra/database/repositories"

	"log"
)

type TransactionTypeRepository struct{}

type TransactionTypeRepositoryInterface interface {
	FindByColumn(tx *sql.Tx, conn *sql.Conn, column string, value any) (t transactionType.TransactionType)
}

func (*TransactionTypeRepository) FindByColumn(tx *sql.Tx, conn *sql.Conn, colunm string, value any) (t transactionType.TransactionType) {
	query := `SELECT transaction_type, description FROM dock_account.transaction_type WHERE ` + colunm + ` = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("TTRFBC 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(value).Scan(&t.TransactionType, &t.Description)

	if err != nil {
		log.Println("TTRFBC 02: ", err)
	}

	return
}
