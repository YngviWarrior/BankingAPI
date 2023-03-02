package mysql

import (
	accountStatement "api-go/core/entities/accountStatement"
	"api-go/infra/database/repositories"
	"api-go/infra/utils"
	"database/sql"

	"log"
)

type AccountStatementRepository struct{}

type AccountStatementRepositoryInterface interface {
	List(tx *sql.Tx, conn *sql.Conn, accountId uint64, dateStart, dateFinish string) (list []*accountStatement.AccountStatementTransactionType)
	Create(tx *sql.Tx, conn *sql.Conn, u accountStatement.AccountStatement) bool
	UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool
}

func (*AccountStatementRepository) UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool {
	_, wheres, updates := utils.QueryFormatter(updateFields, updatefieldValues, wherecolumns, wherevalues, paginationValues, order)
	query := `UPDATE account SET ` + updates + wheres

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("ASRUD 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = tx.Exec(query)

	if err != nil {
		log.Println("ASRUD 02: ", err)
		return false
	}

	return true
}

func (h *AccountStatementRepository) Create(tx *sql.Tx, conn *sql.Conn, acc accountStatement.AccountStatement) bool {
	query := `INSERT INTO account_statement (account, transaction_type, previous_balance, current_balance) VALUES (?, ?, ?, ?)`

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

func (*AccountStatementRepository) List(tx *sql.Tx, conn *sql.Conn, accountId uint64, dateStart, dateFinish string) (list []*accountStatement.AccountStatementTransactionType) {
	query := `
	SELECT a.account_statement, a.account, a.transaction_type, a.previous_balance, a.current_balance, a.registered_at, tp.description
	FROM account_statement a
	JOIN transaction_type tp ON a.transaction_type = tp.transaction_type
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
