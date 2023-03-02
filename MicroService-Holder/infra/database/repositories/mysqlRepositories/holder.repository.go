package mysql

import (
	"database/sql"
	holder "holder-ms/core/entities/holder"
	"holder-ms/infra/database/repositories"
	"holder-ms/infra/utils"

	"log"
)

type HolderRepository struct{}

type HolderRepositoryInterface interface {
	FindByColumn(tx *sql.Tx, conn *sql.Conn, column string, value any) (u holder.Holder)
	Create(tx *sql.Tx, conn *sql.Conn, u holder.Holder) bool
	Verify(tx *sql.Tx, conn *sql.Conn, holderId uint64) bool
	UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool
}

func (*HolderRepository) UpdateDynamically(tx *sql.Tx, conn *sql.Conn, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool {
	_, wheres, updates := utils.QueryFormatter(updateFields, updatefieldValues, wherecolumns, wherevalues, paginationValues, order)
	query := `UPDATE dock_holder.holder SET ` + updates + wheres

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("HRUD 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		log.Println("HRUD 02: ", err)
		return false
	}

	return true
}

func (*HolderRepository) Verify(tx *sql.Tx, conn *sql.Conn, holderId uint64) bool {
	query := `UPDATE dock_holder.holder SET verified = 1 WHERE holder = ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("HRV 01: ", err)
		return false
	}

	defer stmt.Close()

	res, err := stmt.Exec(holderId)

	if err != nil {
		log.Panic("HRV 02: ", err)
		return false
	}

	affcRows, _ := res.RowsAffected()

	if affcRows == 0 {
		log.Println("HRV 03")
		return false
	}

	return true
}

func (h *HolderRepository) Create(tx *sql.Tx, conn *sql.Conn, holder holder.Holder) bool {
	query := `INSERT INTO dock_holder.holder(full_name, cpf) VALUES(?, ?)`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panic("HRC 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(&holder.FullName, &holder.CPF)

	if err != nil {
		log.Panic("HRC 02: ", err)
		return false
	}

	return true
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
