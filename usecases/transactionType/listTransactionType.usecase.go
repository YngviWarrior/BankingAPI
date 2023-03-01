package transactionTypeUseCase

import (
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
)

type ListTransactionTypeUsecase struct {
	Database                  database.DatabaseInterface
	HolderRepository          repository.HolderRepositoryInterface
	TransactionTypeRepository repository.TransactionTypeRepositoryInterface
}

func (c *ListTransactionTypeUsecase) ListTransactionType(input *InputListTransactionTypeDto) (output OutputListTransactionTypeDto, err error) {
	conn := c.Database.CreateConnection()

	typelist := c.TransactionTypeRepository.List(nil, conn)

	for _, ty := range typelist {
		var t transactionType = transactionType(*ty)

		output.List = append(output.List, &t)
	}

	conn.Close()
	return
}
