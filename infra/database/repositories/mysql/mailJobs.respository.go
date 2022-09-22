package mysql

import (
	constants "api-user/core"
	database "api-user/infra/database"
	"log"
	"time"
)

type MailJobsRepository struct{}

type MailJobsRepositoryInterface interface {
	CreateJob(userId uint64, layoutId int64, ref2 int64, ref3 int64, ref4 int64, ref5 int64, ref6 int64, ref7 int64, sendDate string) bool
}

func (*MailJobsRepository) CreateJob(userId uint64, layoutId int64, ref2 int64, ref3 int64, ref4 int64, ref5 int64, ref6 int64, ref7 int64, sendDate string) bool {
	conn := database.GetConnection()

	if sendDate == "" {
		sendDate = time.Now().Format("2006-01-02 15:04:05")
	}

	_, err := conn.Exec(`
		INSERT INTO tarefas (id_tipo, id_status, id_usuario, data_registro, data_proxima_tentativa, ref1, ref2, ref3, ref4, ref5, ref6, ref7) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, constants.TAREFAS_ID_TIPO_SEND_EMAIL, constants.TAREFAS_ID_STATUS_PENDENTE, userId, time.Now().Format("2006-01-02 15:04:05"), sendDate, layoutId, ref2, ref3, ref4, ref5, ref6, ref7)

	if err != nil {
		log.Println("CJMJR 01: ", err)
		return false
	}

	defer conn.Close()

	return true
}
