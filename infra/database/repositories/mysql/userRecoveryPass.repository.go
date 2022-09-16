package mysql

import (
	database "go-api/infra/database"
	"log"
)

type UserRecoveryPassRepository struct{}

type UserRecoveryPassRepositoryInterface interface {
	CreateRecoveryPass(userId uint64, code string, registerTime string, expireTime string) int64
}

func (*UserRecoveryPassRepository) CreateRecoveryPass(userId uint64, code string, registerTime string, expireTime string) int64 {
	conn := database.GetConnection()

	res, err := conn.Exec(`
		INSERT INTO usuarios_recuperar_senha(id_usuario, code, nova_senha_hash, data_alterado, data_registro, data_expira) 
		VALUES (?, ?, ?, ?, ?, ?)
	`, userId, code, nil, nil, registerTime, expireTime)

	if err != nil {
		log.Println("URPR 01: ", err)
		return 0
	}

	defer conn.Close()

	lastInsert, _ := res.LastInsertId()

	return lastInsert
}
