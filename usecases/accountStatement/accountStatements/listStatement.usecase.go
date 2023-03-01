package listStatementUseCase

import (
	accountEntity "api-go/core/entities/account"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	"fmt"
)

type ListStatementUsecase struct {
	Database                   database.DatabaseInterface
	AccountStatementRepository repository.AccountStatementRepositoryInterface
	AccountRepository          repository.AccountRepositoryInterface
}

func (c *ListStatementUsecase) ListStatement(input *InputListStatementDto) (output OutputListStatementDto, err error) {
	conn := c.Database.CreateConnection()

	acc := c.AccountRepository.FindByColumn(nil, conn, "account", input.Account)

	if (acc == accountEntity.Account{}) {
		err = fmt.Errorf("account dosent exists")
		return
	}

	statementList := c.AccountStatementRepository.List(nil, conn, acc.Account, input.DateStart, input.DateFinish)

	if len(statementList) == 0 {
		return
	}

	for _, v := range statementList {
		output.List = append(output.List, *v)
	}

	conn.Close()
	return
}
