package mysql

import (
	account "account-ms/core/entities/account"
	"account-ms/infra/database/repositories"
	"account-ms/infra/utils"
	"database/sql"
	"time"

	"log"
)

type AccountRepository struct{}

type AccountRepositoryInterface interface {
	Find(tx *sql.Tx, conn *sql.Conn, agency string, number any) (u account.AccountHolder)
	FindByColumn(tx *sql.Tx, conn *sql.Conn, column string, value any) (u account.Account)
	Create(tx *sql.Tx, conn *sql.Conn, u account.Account) bool
	UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool
}

func (*AccountRepository) UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool {
	_, wheres, updates := utils.QueryFormatter(updateFields, updatefieldValues, wherecolumns, wherevalues, paginationValues, order)
	query := `UPDATE dock_account.account SET ` + updates + wheres

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

func (h *AccountRepository) Create(tx *sql.Tx, conn *sql.Conn, acc account.Account) bool {
	query := `CALL dock_account.sp_createAccount(?, ?, ?, ?)`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("ARC 01: ", err)
		return false
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(&acc.Holder, &acc.Agency, &acc.Number, &date)

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
		h.full_name, h.cpf, h.verified, h.activated as holder_activated
	FROM dock_account.account a 
	JOIN dock_holder.holder h ON a.holder = h.holder
	WHERE a.agency = ? AND a.number = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("ARF 01: ", err)
		return
	}

	err = stmt.QueryRow(agency, number).Scan(&u.Account, &u.Holder, &u.Agency, &u.Number, &u.Balance, &u.Activated, &u.Blocked, &u.HolderName, &u.HolderDoc, &u.HolderVerified, &u.HolderActivated)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panic("ARF 02: ", err)
		return
	}

	return
}

func (*AccountRepository) FindByColumn(tx *sql.Tx, conn *sql.Conn, colunm string, value any) (u account.Account) {
	query := `SELECT account, holder, agency, number, balance, activated, blocked FROM dock_account.account WHERE ` + colunm + ` = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("ARFBC 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(value).Scan(&u.Account, &u.Holder, &u.Agency, &u.Number, &u.Balance, &u.Activated, &u.Blocked)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panic("ARFBC 02: ", err)
		return
	}

	return
}
