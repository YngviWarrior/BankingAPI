package controllers

import (
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	usecase "api-go/usecases/holder/holderCreate"
	"encoding/json"
	"log"
	"net/http"
)

type inputCreateHoldersDto struct {
	FullName *string `json:"full_name" validate:"required"`
	CPF      *string `json:"cpf" validate:"cpf,required"`
}

func (c *Controllers) HandlerCreateHolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputCreateHoldersDto
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

	useCaseInputDto := usecase.InputCreateHolderDto{FullName: *input.FullName, CPF: *input.CPF}

	var dbInterface database.DatabaseInterface = c.DataBase
	var holderUserInterface repository.HolderRepositoryInterface = &repository.HolderRepository{}
	var useCase = usecase.CreateHolderUsecase{}

	useCase.Database = dbInterface
	useCase.HolderRepository = holderUserInterface

	_, err = useCase.CreateHolder(&useCaseInputDto)

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
