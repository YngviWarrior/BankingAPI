package mysql

import (
	holder "account-ms/core/entities/holder"
	"account-ms/infra/database/repositories"
	"database/sql"

	"log"
)

type HolderRepository struct{}

type HolderRepositoryInterface interface {
	FindByColumn(tx *sql.Tx, conn *sql.Conn, column string, value any) (u holder.Holder)
}

func (*HolderRepository) FindByColumn(tx *sql.Tx, conn *sql.Conn, colunm string, value any) (u holder.Holder) {
	query := `SELECT holder, full_name, cpf, verified, activated FROM dock_holder.holder WHERE ` + colunm + ` = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Println("HRFBC 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(value).Scan(&u.Holder, &u.FullName, &u.CPF, &u.Verified, &u.Activated)

	if err != nil {
		log.Println("HRFBC 02: ", err)
	}

	return
}
