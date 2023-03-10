package controllers

import (
	"encoding/json"
	"holder-ms/infra/database"
	repository "holder-ms/infra/database/repositories/mysqlRepositories"
	usecase "holder-ms/usecases/holder/holderFind"
	"log"
	"net/http"
)

type inputFindHoldersDto struct {
	CPF *string `json:"cpf" validate:"required"`
}

func (c *Controllers) HandlerFindHolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputFindHoldersDto
	var send outputControllerDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Printf("SI01: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		send.Errors = append(send.Errors, "invalid primitive type")
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

	useCaseInputDto := usecase.InputFindHolderDto{CPF: *input.CPF}

	var dbInterface database.DatabaseInterface = c.DataBase
	var holderUserInterface repository.HolderRepositoryInterface = &repository.HolderRepository{}
	var accountInterface repository.AccountRepositoryInterface = &repository.AccountRepository{}
	var useCase = usecase.FindHolderUsecase{}

	useCase.Database = dbInterface
	useCase.HolderRepository = holderUserInterface
	useCase.AccountRepository = accountInterface

	output, err := useCase.FindHolder(&useCaseInputDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else {
		send.Status = 1
		send.Message = "Success"
		send.Data = output
	}

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Panicf("SI02: %s", err)
	}

	w.Write(jsonResp)
}
