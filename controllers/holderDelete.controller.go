package controllers

import (
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	usecase "api-go/usecases/holder/holderDelete"
	"encoding/json"
	"log"
	"net/http"
)

type inputDeleteHoldersDto struct {
	CPF *string `json:"cpf" validate:"required"`
}

func (c *Controllers) HandlerDeleteHolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputDeleteHoldersDto
	var send outputControllerDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Printf("SI01: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		send.Errors = append(send.Errors, "body is invalid")
		jsonResp, err := json.Marshal(send)

		if err != nil {
			log.Fatalf("SI02: %s", err)
		}

		w.Write(jsonResp)
		return
	}

	if !c.InputValidation(w, input) {
		return
	}

	useCaseInputDto := usecase.InputDeleteHolderDto{CPF: *input.CPF}

	var dbInterface database.DatabaseInterface = c.DataBase
	var holderUserInterface repository.HolderRepositoryInterface = &repository.HolderRepository{}
	var transactionTypeRepoInterface repository.TransactionTypeRepositoryInterface = &repository.TransactionTypeRepository{}
	var accountRepoInterface repository.AccountRepositoryInterface = &repository.AccountRepository{}
	var accountStatementRepoInterface repository.AccountStatementRepositoryInterface = &repository.AccountStatementRepository{}
	var useCase = usecase.DeleteHolderUsecase{}

	useCase.Database = dbInterface
	useCase.HolderRepository = holderUserInterface
	useCase.TransactionTypeRepository = transactionTypeRepoInterface
	useCase.AccountRepository = accountRepoInterface
	useCase.AccountStatementRepository = accountStatementRepoInterface

	_, err = useCase.DeleteHolder(&useCaseInputDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else {
		send.Status = 1
		send.Message = "Success"
		// send.Data = output
	}

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Panicf("SI02: %s", err)
	}

	w.Write(jsonResp)
}
