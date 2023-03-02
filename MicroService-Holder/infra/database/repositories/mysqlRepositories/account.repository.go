package mysql

import (
	"database/sql"
	account "holder-ms/core/entities/account"
	"holder-ms/infra/database/repositories"
	"holder-ms/infra/utils"

	"log"
)

type AccountRepository struct{}

type AccountRepositoryInterface interface {
	ListByColumn(tx *sql.Tx, conn *sql.Conn, column string, value any) (list []*account.Account)
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

func (*AccountRepository) ListByColumn(tx *sql.Tx, conn *sql.Conn, colunm string, value any) (list []*account.Account) {
	query := `SELECT account, holder, agency, number, balance, activated, blocked FROM dock_account.account WHERE ` + colunm + ` = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("ARLBC 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(value)

	if err != nil {
		log.Panic("ARLBC 02: ", err)
		return
	}

	for res.Next() {
		var u account.Account
		err := res.Scan(&u.Account, &u.Holder, &u.Agency, &u.Number, &u.Balance, &u.Activated, &u.Blocked)

		if err != nil {
			log.Panic("ARLBC 03: ", err)
			return
		}

		list = append(list, &u)
	}

	return
}
