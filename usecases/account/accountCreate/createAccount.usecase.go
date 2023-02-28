package createAccountUseCase

import (
	accountEntity "api-go/core/entities/account"
	holderEntity "api-go/core/entities/holder"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	"api-go/infra/utils"
	"database/sql"
	"fmt"
	"log"
)

type CreateAccountUsecase struct {
	Database          database.DatabaseInterface
	HolderRepository  repository.HolderRepositoryInterface
	AccountRepository repository.AccountRepositoryInterface
}

func (c *CreateAccountUsecase) validAccountNumber(conn *sql.Conn, number int64) int64 {
	a := c.AccountRepository.FindByColumn(nil, conn, "number", number)

	if (a != accountEntity.Account{}) {
		return c.validAccountNumber(conn, utils.RandomNumber())
	}

	return number
}

func (c *CreateAccountUsecase) CreateAccount(input *InputCreateAccountDto) (output OutputCreateAccountDto, err error) {
	conn := c.Database.CreateConnection()

	h := c.HolderRepository.FindByColumn(nil, conn, "cpf", input.CPF)

	if (h == holderEntity.Holder{}) {
		err = fmt.Errorf("holder dosent exists")
		conn.Close()
		return
	}

	if !h.Verified {
		err = fmt.Errorf("holder is not verified")
		conn.Close()
		return
	}

	var a accountEntity.Account

	a.Holder = h.Holder
	a.Agency = "0001"
	a.Number = fmt.Sprintf("%d", c.validAccountNumber(conn, utils.RandomNumber()))[:8]

	if !c.AccountRepository.Create(nil, conn, a) {
		log.Println("CACC 01")
		conn.Close()
		return
	}

	conn.Close()
	return
}
