package mysql

import (
	account "api-go/core/entities/account"
	"api-go/infra/database/repositories"
	"api-go/infra/utils"
	"database/sql"

	"log"
)

type AccountRepository struct{}

type AccountRepositoryInterface interface {
	Find(tx *sql.Tx, conn *sql.Conn, agency string, number any) (u account.AccountHolder)
	FindByColumn(tx *sql.Tx, conn *sql.Conn, column string, value any) (u account.Account)
	Create(tx *sql.Tx, conn *sql.Conn, u account.Account) bool
	Delete(tx *sql.Tx, conn *sql.Conn, accountId uint64) bool
	UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool
}

func (*AccountRepository) UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool {
	_, wheres, updates := utils.QueryFormatter(updateFields, updatefieldValues, wherecolumns, wherevalues, paginationValues, order)
	query := `UPDATE account SET ` + updates + wheres

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("ARUD 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		log.Println("ARUD 02: ", err)
		return false
	}

	return true
}

func (h *AccountRepository) Delete(tx *sql.Tx, conn *sql.Conn, AccountId uint64) bool {
	query := `DELETE FROM account WHERE account = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("ARD 01: ", err)
		return false
	}

	_, err = stmt.Exec(AccountId)

	if err != nil {
		log.Panic("ARD 02: ", err)
		return false
	}

	defer stmt.Close()

	return true
}

func (h *AccountRepository) Create(tx *sql.Tx, conn *sql.Conn, acc account.Account) bool {
	query := `CALL sp_createAccount(?, ?, ?)`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("ARC 01: ", err)
		return false
	}

	_, err = stmt.Exec(&acc.Holder, &acc.Agency, &acc.Number)

	if err != nil {
		log.Panic("ARC 02: ", err)
		return false
	}

	defer stmt.Close()

	return true
}

func (*AccountRepository) Find(tx *sql.Tx, conn *sql.Conn, agency string, number any) (u account.AccountHolder) {
	query := `
	SELECT a.account, a.holder, a.agency, a.number, a.balance, a.activated, a.blocked,
		h.full_name, h.cpf
	FROM account a 
	JOIN holder h ON a.holder = h.holder
	WHERE a.agency = ? AND a.number = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("ARF 01: ", err)
		return
	}

	err = stmt.QueryRow(agency, number).Scan(&u.Account, &u.Holder, &u.Agency, &u.Number, &u.Balance, &u.Activated, &u.Blocked, &u.HolderName, &u.HolderDoc)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panic("ARF 02: ", err)
		return
	}

	return
}

func (*AccountRepository) FindByColumn(tx *sql.Tx, conn *sql.Conn, colunm string, value any) (u account.Account) {
	query := `SELECT account, holder, agency, number, balance, activated, blocked FROM account WHERE ` + colunm + ` = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("ARFBC 01: ", err)
		return
	}

	err = stmt.QueryRow(value).Scan(&u.Account, &u.Holder, &u.Agency, &u.Number, &u.Balance, &u.Activated, &u.Blocked)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panic("ARFBC 02: ", err)
		return
	}

	return
}
