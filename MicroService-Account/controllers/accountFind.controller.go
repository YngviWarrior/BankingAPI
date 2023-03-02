package controllers

import (
	"account-ms/infra/database"
	repository "account-ms/infra/database/repositories/mysqlRepositories"
	usecase "account-ms/usecases/account/accountFind"
	"encoding/json"
	"log"
	"net/http"
)

type inputFindAccountDto struct {
	Agency *string `json:"agency" validate:"required"`
	Number *string `json:"number" validate:"required"`
}

func (c *Controllers) HandlerFindAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputFindAccountDto
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

	useCaseInputDto := usecase.InputFindAccountDto{Agency: *input.Agency, Number: *input.Number}

	var dbInterface database.DatabaseInterface = c.DataBase
	var accountInterface repository.AccountRepositoryInterface = &repository.AccountRepository{}
	var useCase = usecase.FindAccountUseCase{}

	useCase.Database = dbInterface
	useCase.AccountRepository = accountInterface

	output, err := useCase.FindAccount(&useCaseInputDto)

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
