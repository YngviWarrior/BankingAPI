package transactionAccountUseCase

import (
	accountEntity "account-ms/core/entities/account"
	accountStatementEntity "account-ms/core/entities/accountStatement"
	transactionTypeEntity "account-ms/core/entities/transactionType"
	"account-ms/infra/database"
	repository "account-ms/infra/database/repositories/mysqlRepositories"
	"account-ms/infra/utils"
	"fmt"
	"time"
)

type TransactionAccountUsecase struct {
	Database                   database.DatabaseInterface
	TransactionTypeRepository  repository.TransactionTypeRepositoryInterface
	AccountRepository          repository.AccountRepositoryInterface
	AccountStatementRepository repository.AccountStatementRepositoryInterface
}

var WITHDRAW_THRESHOULD float64 = 2000.00

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

	if !a.HolderVerified {
		err = fmt.Errorf("holder isnt verified")
		conn.Close()
		return
	}

	if !a.HolderActivated {
		err = fmt.Errorf("holder is deactivated")
		conn.Close()
		return
	}

	if a.Blocked {
		err = fmt.Errorf("account is blocked")
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
		s.RegisteredDate = time.Now().Format("2006-01-02 15:04:05")

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
		tStart := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
		tEnd := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 59, time.Local).Format("2006-01-02 15:04:05")
		todayStatements := c.AccountStatementRepository.List(nil, conn, a.Account, tStart, tEnd)

		var dailyWithdraw float64 = 0
		for _, statement := range todayStatements {
			dailyWithdraw += statement.PreviousBalance - statement.CurrentBalance
		}

		if dailyWithdraw >= WITHDRAW_THRESHOULD {
			err = fmt.Errorf("your withdraw limit is already reached")
			conn.Close()
			return
		} else if (dailyWithdraw+input.Amount) > WITHDRAW_THRESHOULD || input.Amount > WITHDRAW_THRESHOULD {
			err = fmt.Errorf("your withdraw limit is (%.f), you can only withdraw (%.f)", utils.ToFixed(WITHDRAW_THRESHOULD, 2), utils.ToFixed(WITHDRAW_THRESHOULD-dailyWithdraw, 2))
			conn.Close()
			return
		}

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
