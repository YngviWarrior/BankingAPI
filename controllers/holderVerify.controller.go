package controllers

import (
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	usecase "api-go/usecases/holder/holderVerify"
	"encoding/json"
	"log"
	"net/http"
)

type inputVerifyHoldersDto struct {
	CPF *string `json:"cpf" validate:"required"`
}

func (c *Controllers) HandlerVerifyHolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputVerifyHoldersDto
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

	useCaseInputDto := usecase.InputVerifyHolderDto{CPF: *input.CPF}

	var dbInterface database.DatabaseInterface = c.DataBase
	var holderUserInterface repository.HolderRepositoryInterface = &repository.HolderRepository{}
	var useCase = usecase.VerifyHolderUsecase{}

	useCase.Database = dbInterface
	useCase.HolderRepository = holderUserInterface

	_, err = useCase.VerifyHolder(&useCaseInputDto)

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
