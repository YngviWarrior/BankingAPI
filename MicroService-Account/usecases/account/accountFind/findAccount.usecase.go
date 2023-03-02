package findAccountUseCase

import (
	accountEntity "account-ms/core/entities/account"
	"account-ms/infra/database"
	repository "account-ms/infra/database/repositories/mysqlRepositories"
	"fmt"
)

type FindAccountUseCase struct {
	Database          database.DatabaseInterface
	AccountRepository repository.AccountRepositoryInterface
}

func (c *FindAccountUseCase) FindAccount(input *InputFindAccountDto) (output OutputFindAccountDto, err error) {
	conn := c.Database.CreateConnection()

	a := c.AccountRepository.Find(nil, conn, input.Agency, input.Number)

	if (a == accountEntity.AccountHolder{}) {
		conn.Close()
		err = fmt.Errorf("account dont exists")
		return
	}

	output.HolderName = a.HolderName
	output.HolderDoc = a.HolderDoc
	output.Agency = a.Agency
	output.Number = a.Number
	output.Balance = a.Balance
	output.Activated = a.Activated
	output.Blocked = a.Blocked

	conn.Close()
	return
}
