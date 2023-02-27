package transactionAccountUseCase

import (
	accountEntity "api-go/core/entities/account"
	accountStatementEntity "api-go/core/entities/accountStatement"
	transactionTypeEntity "api-go/core/entities/transactionType"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	"api-go/infra/utils"
	"fmt"
)

type TransactionAccountUsecase struct {
	Database                   database.DatabaseInterface
	TransactionTypeRepository  repository.TransactionTypeRepositoryInterface
	AccountRepository          repository.AccountRepositoryInterface
	AccountStatementRepository repository.AccountStatementRepositoryInterface
}

func (c *TransactionAccountUsecase) TransactionAccount(input *InputTransactionAccountDto) (output *OutputTransactionAccountDto, err error) {
	conn := c.Database.CreateConnection()

	t := c.TransactionTypeRepository.FindByColumn(nil, conn, "transaction_type", input.TransactionType)

	if (t == transactionTypeEntity.TransactionType{}) {
		err = fmt.Errorf("invalid transaction type")
		conn.Close()
		return
	}

	switch t.TransactionType {
	case 1, 4:
		err = fmt.Errorf("invalid transaction type")
		conn.Close()
		return
	}

	a := c.AccountRepository.Find(nil, conn, input.Agency, input.Number)

	if (a == accountEntity.AccountHolder{}) {
		err = fmt.Errorf("account dosent exists")
		conn.Close()
		return
	}

	switch t.TransactionType {
	case 2: // deposit
		var s accountStatementEntity.AccountStatement

		s.Account = a.Account
		s.PreviousBalance = a.Balance
		s.CurrentBalance = a.Balance + utils.ToFixed(input.Amount, 2)
		s.TransactionType = t.TransactionType

		if !c.AccountStatementRepository.Create(nil, conn, s) {
			err = fmt.Errorf("cant add statement")
			conn.Close()
			return
		}

		if !c.AccountRepository.UpdateDynamically(nil, conn, []string{"balance"}, []any{utils.ToFixed(s.CurrentBalance, 2)}, []string{"agency", "number"}, []any{a.Agency, a.Number}, []any{}, "") {
			err = fmt.Errorf("cant change balance")
			conn.Close()
			return
		}
	case 3: // withdraw
		if a.Balance < input.Amount {
			err = fmt.Errorf("insufficient balance: your current balance is (%.f)", utils.ToFixed(a.Balance, 2))
			conn.Close()
			return
		}

		var s accountStatementEntity.AccountStatement

		s.Account = a.Account
		s.PreviousBalance = a.Balance
		s.CurrentBalance = a.Balance - utils.ToFixed(input.Amount, 2)
		s.TransactionType = t.TransactionType

		if !c.AccountStatementRepository.Create(nil, conn, s) {
			err = fmt.Errorf("cant add statement")
			conn.Close()
			return
		}

		if !c.AccountRepository.UpdateDynamically(nil, conn, []string{"balance"}, []any{utils.ToFixed(s.CurrentBalance, 2)}, []string{"agency", "number"}, []any{a.Agency, a.Number}, []any{}, "") {
			err = fmt.Errorf("cant change balance")
			conn.Close()
			return
		}
	}

	conn.Close()
	return
}
