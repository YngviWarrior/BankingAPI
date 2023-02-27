package mysql

import (
	transactionType "api-go/core/entities/transactionType"
	"api-go/infra/database/repositories"
	"database/sql"

	"log"
)

type TransactionTypeRepository struct{}

type TransactionTypeRepositoryInterface interface {
	List(tx *sql.Tx, conn *sql.Conn) (list []*transactionType.TransactionType)
	FindByColumn(tx *sql.Tx, conn *sql.Conn, column string, value any) (t transactionType.TransactionType)
}

func (*TransactionTypeRepository) FindByColumn(tx *sql.Tx, conn *sql.Conn, colunm string, value any) (t transactionType.TransactionType) {
	query := `SELECT transaction_type, description FROM transaction_type WHERE ` + colunm + ` = ?`

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

func (*TransactionTypeRepository) List(tx *sql.Tx, conn *sql.Conn) (list []*transactionType.TransactionType) {
	query := `
	SELECT transaction_type, description
	FROM transaction_type`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("TTRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panic("TTRL 02: ", err)
		return
	}

	for res.Next() {
		var a transactionType.TransactionType

		err := res.Scan(&a.TransactionType, &a.Description)

		if err != nil {
			log.Panic("TTRL 03: ", err)
			return
		}

		list = append(list, &a)
	}

	return
}
