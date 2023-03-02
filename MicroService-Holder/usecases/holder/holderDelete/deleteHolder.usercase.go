package deleteHolderUseCase

import (
	"context"
	"fmt"
	accountStatementEntity "holder-ms/core/entities/accountStatement"
	holderEntity "holder-ms/core/entities/holder"
	"holder-ms/infra/database"
	repository "holder-ms/infra/database/repositories/mysqlRepositories"
)

type DeleteHolderUsecase struct {
	Database                   database.DatabaseInterface
	HolderRepository           repository.HolderRepositoryInterface
	TransactionTypeRepository  repository.TransactionTypeRepositoryInterface
	AccountRepository          repository.AccountRepositoryInterface
	AccountStatementRepository repository.AccountStatementRepositoryInterface
}

func (c *DeleteHolderUsecase) DeleteHolder(input *InputDeleteHolderDto) (output OutputDeleteHolderDto, err error) {
	ctx := context.TODO()
	conn := c.Database.CreateConnection()

	h := c.HolderRepository.FindByColumn(nil, conn, "cpf", input.CPF)

	if (h == holderEntity.Holder{}) {
		err = fmt.Errorf("holder dosent exists")
		conn.Close()
		return
	}

	t := c.TransactionTypeRepository.FindByColumn(nil, conn, "description", "withdraw")

	accList := c.AccountRepository.ListByColumn(nil, conn, "holder", h.Holder)

	if len(accList) != 0 {
		tx := c.Database.CreateTransaction(&ctx, conn)

		for _, acc := range accList {
			if acc.Balance > 0 {
				if !c.AccountStatementRepository.Create(tx, nil, accountStatementEntity.AccountStatement{
					Account:         acc.Account,
					TransactionType: t.TransactionType,
					PreviousBalance: acc.Balance,
					CurrentBalance:  0,
				}) {
					err = fmt.Errorf("failed to clean accounts")
					tx.Rollback()
					conn.Close()
					return
				}
			}

			if !c.AccountRepository.UpdateDynamically(tx, nil, []string{"activated"}, []any{0}, []string{"account"}, []any{acc.Account}, []any{}, "") {
				err = fmt.Errorf("failed deactive an account")
				tx.Rollback()
				conn.Close()
				return
			}
		}

		tx.Commit()
	}

	if !c.HolderRepository.UpdateDynamically(nil, conn, []string{"activated"}, []any{0}, []string{"holder"}, []any{h.Holder}, []any{}, "") {
		err = fmt.Errorf("failed to deactivate holder")
		conn.Close()
		return
	}

	conn.Close()
	return
}
