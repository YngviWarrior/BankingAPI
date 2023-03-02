package findHolderUseCase

import (
	"fmt"
	holderEntity "holder-ms/core/entities/holder"
	"holder-ms/infra/database"
	repository "holder-ms/infra/database/repositories/mysqlRepositories"
)

type FindHolderUsecase struct {
	Database          database.DatabaseInterface
	HolderRepository  repository.HolderRepositoryInterface
	AccountRepository repository.AccountRepositoryInterface
}

func (c *FindHolderUsecase) FindHolder(input *InputFindHolderDto) (output OutputFindHolderDto, err error) {
	conn := c.Database.CreateConnection()

	h := c.HolderRepository.FindByColumn(nil, conn, "cpf", input.CPF)

	if (h == holderEntity.Holder{}) {
		conn.Close()
		err = fmt.Errorf("holder not found")
		return
	}

	output.FullName = h.FullName
	output.CPF = h.CPF
	output.Activated = h.Activated
	output.Verified = h.Verified

	acclist := c.AccountRepository.ListByColumn(nil, conn, "holder", h.Holder)

	for _, acc := range acclist {
		var a account

		a.Activated = acc.Activated
		a.Agency = acc.Agency
		a.Number = acc.Number
		a.Balance = acc.Balance
		a.Blocked = acc.Blocked
		a.HolderDoc = h.CPF
		a.HolderName = h.FullName

		output.AccountList = append(output.AccountList, &a)
	}

	conn.Close()
	return
}
