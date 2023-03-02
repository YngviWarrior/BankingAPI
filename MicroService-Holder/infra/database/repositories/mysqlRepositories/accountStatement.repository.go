package mysql

import (
	"database/sql"
	accountStatement "holder-ms/core/entities/accountStatement"
	"holder-ms/infra/database/repositories"

	"log"
)

type AccountStatementRepository struct{}

type AccountStatementRepositoryInterface interface {
	Create(tx *sql.Tx, conn *sql.Conn, u accountStatement.AccountStatement) bool
}

func (h *AccountStatementRepository) Create(tx *sql.Tx, conn *sql.Conn, acc accountStatement.AccountStatement) bool {
	query := `INSERT INTO dock_account.account_statement (account, transaction_type, previous_balance, current_balance) VALUES (?, ?, ?, ?)`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("ASRC 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(&acc.Account, &acc.TransactionType, &acc.PreviousBalance, &acc.CurrentBalance)

	if err != nil {
		log.Panic("ASRC 02: ", err)
		return false
	}

	return true
}
