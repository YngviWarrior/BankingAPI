package controllers

import (
	"account-ms/infra/database"
	repository "account-ms/infra/database/repositories/mysqlRepositories"
	usecase "account-ms/usecases/accountStatement/accountStatements"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type inputListStatementDto struct {
	AccountAgency *string `json:"account_agency" validate:"required"`
	AccountNumber *string `json:"account_number" validate:"required"`
	DateStart     *string `json:"date_start" validate:"date,required"`
	DateFinish    *string `json:"date_finish" validate:"date,required"`
}

func (c *Controllers) HandlerListStatement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputListStatementDto
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

	useCaseInputDto := usecase.InputListStatementDto{AccountAgency: *input.AccountAgency, AccountNumber: *input.AccountNumber, DateStart: fmt.Sprintf("%s 00:00:00", *input.DateStart), DateFinish: fmt.Sprintf("%s 23:59:59", *input.DateFinish)}

	var dbInterface database.DatabaseInterface = c.DataBase
	var AccountRepoInterface repository.AccountRepositoryInterface = &repository.AccountRepository{}
	var AccountStatementRepoInterface repository.AccountStatementRepositoryInterface = &repository.AccountStatementRepository{}
	var useCase = usecase.ListStatementUsecase{}

	useCase.Database = dbInterface
	useCase.AccountRepository = AccountRepoInterface
	useCase.AccountStatementRepository = AccountStatementRepoInterface

	output, err := useCase.ListStatement(&useCaseInputDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else if len(output.List) == 0 {
		send.Status = 1
		send.Message = "Success"
		// send.Data = output.List
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
