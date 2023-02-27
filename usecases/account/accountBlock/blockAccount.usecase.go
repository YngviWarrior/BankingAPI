package blockAccountUseCase

import (
	accountEntity "api-go/core/entities/account"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	"fmt"
)

type BlockAccountUseCase struct {
	Database          database.DatabaseInterface
	AccountRepository repository.AccountRepositoryInterface
}

func (c *BlockAccountUseCase) BlockAccount(input *InputBlockAccountDto) (output OutputBlockAccountDto, err error) {
	conn := c.Database.CreateConnection()

	a := c.AccountRepository.Find(nil, conn, input.Agency, input.Number)

	if (a == accountEntity.AccountHolder{}) {
		err = fmt.Errorf("account dosent exists")
		conn.Close()
		return
	}

	if !c.AccountRepository.UpdateDynamically(nil, conn, []string{"blocked"}, []any{input.Block}, []string{"agency", "number"}, []any{a.Agency, a.Number}, []any{}, "") {
		err = fmt.Errorf("account dont exists")
		conn.Close()
		return
	}

	conn.Close()
	return
}
