package mysql

import (
	accountStatement "account-ms/core/entities/accountStatement"
	"account-ms/infra/database/repositories"
	"database/sql"

	"log"
)

type AccountStatementRepository struct{}

type AccountStatementRepositoryInterface interface {
	List(tx *sql.Tx, conn *sql.Conn, accountId uint64, dateStart, dateFinish string) (list []*accountStatement.AccountStatementTransactionType)
	Create(tx *sql.Tx, conn *sql.Conn, u accountStatement.AccountStatement) bool
}

func (h *AccountStatementRepository) Create(tx *sql.Tx, conn *sql.Conn, acc accountStatement.AccountStatement) bool {
	query := `INSERT INTO dock_account.account_statement (account, transaction_type, previous_balance, current_balance, registered_at) VALUES (?, ?, ?, ?, ?)`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("ASRC 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(&acc.Account, &acc.TransactionType, &acc.PreviousBalance, &acc.CurrentBalance, &acc.RegisteredDate)

	if err != nil {
		log.Panic("ASRC 02: ", err)
		return false
	}

	return true
}

func (*AccountStatementRepository) List(tx *sql.Tx, conn *sql.Conn, accountId uint64, dateStart, dateFinish string) (list []*accountStatement.AccountStatementTransactionType) {
	query := `
	SELECT a.account_statement, a.account, a.transaction_type, a.previous_balance, a.current_balance, a.registered_at, tp.description
	FROM dock_account.account_statement a
	JOIN dock_account.transaction_type tp ON a.transaction_type = tp.transaction_type
	WHERE a.account = ? AND a.registered_at BETWEEN "` + dateStart + `" AND "` + dateFinish + `"`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("ASRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(accountId)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panic("ASRL 02: ", err)
		return
	}

	for res.Next() {
		var a accountStatement.AccountStatementTransactionType

		err := res.Scan(&a.AccountStatement, &a.Account, &a.TransactionType, &a.PreviousBalance, &a.CurrentBalance, &a.RegisteredDate, &a.TransactionTypeDescription)

		if err != nil {
			log.Panic("ASRL 03: ", err)
			return
		}

		list = append(list, &a)
	}

	return
}
