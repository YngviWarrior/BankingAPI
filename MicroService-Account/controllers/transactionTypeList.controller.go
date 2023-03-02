package controllers

import (
	"account-ms/infra/database"
	repository "account-ms/infra/database/repositories/mysqlRepositories"
	usecase "account-ms/usecases/transactionType"
	"encoding/json"
	"log"
	"net/http"
)

// type inputListTransactionTypeDto struct{}

func (c *Controllers) HandlerListTransactionType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var send outputControllerDto

	useCaseInputDto := usecase.InputListTransactionTypeDto{}

	var dbInterface database.DatabaseInterface = c.DataBase
	var transactionTypeInterface repository.TransactionTypeRepositoryInterface = &repository.TransactionTypeRepository{}
	var useCase = usecase.ListTransactionTypeUsecase{}

	useCase.Database = dbInterface
	useCase.TransactionTypeRepository = transactionTypeInterface

	output, err := useCase.ListTransactionType(&useCaseInputDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else {
		send.Status = 1
		send.Message = "Success"
		send.Data = output.List
	}

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Panicf("SI02: %s", err)
	}

	w.Write(jsonResp)
}
