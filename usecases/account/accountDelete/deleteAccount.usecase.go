package deleteAccountUseCase

import (
	accountEntity "api-go/core/entities/account"
	accountStatementEntity "api-go/core/entities/accountStatement"
	transactionTypeEntity "api-go/core/entities/transactionType"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	"context"
	"fmt"
)

type DeleteAccountUsecase struct {
	Database                   database.DatabaseInterface
	TransactionTypeRepository  repository.TransactionTypeRepositoryInterface
	AccountRepository          repository.AccountRepositoryInterface
	AccountStatementRepository repository.AccountStatementRepositoryInterface
}

func (c *DeleteAccountUsecase) DeleteAccount(input *InputDeleteAccountDto) (output *OutputDeleteAccountDto, err error) {
	ctx := context.TODO()
	conn := c.Database.CreateConnection()
	tx := c.Database.CreateTransaction(&ctx, conn)

	a := c.AccountRepository.Find(tx, nil, fmt.Sprintf("%v", input.Agency), input.Number)

	if (a == accountEntity.AccountHolder{}) {
		err = fmt.Errorf("account dosent exists")
		conn.Close()
		return
	}

	t := c.TransactionTypeRepository.FindByColumn(tx, nil, "description", "withdraw")

	if (t == transactionTypeEntity.TransactionType{}) {
		err = fmt.Errorf("transaction_type dosent exists")
		conn.Close()
		return
	}

	if a.Balance > 0 {
		var stmt accountStatementEntity.AccountStatement

		stmt.Account = a.Account
		stmt.TransactionType = t.TransactionType
		stmt.CurrentBalance = 0
		stmt.PreviousBalance = a.Balance

		if !c.AccountStatementRepository.Create(tx, nil, stmt) {
			err = fmt.Errorf("cant create account ending statement")
			conn.Close()
			return
		}
	}

	if !c.AccountRepository.UpdateDynamically(tx, nil, []string{"activated", "balance"}, []any{false, 0}, []string{"agency", "account"}, []any{a.Agency, a.Account}, []any{}, "") {
		err = fmt.Errorf("cant desactivate account")
		conn.Close()
		return
	}

	tx.Commit()
	conn.Close()
	return
}
