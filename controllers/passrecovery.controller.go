package controllers

import (
	"encoding/json"
	repository "go-api/infra/database/repositories/mysql"
	usecase "go-api/usecases/passrecovery"
	"log"
	"net/http"
)

type inputPassRecoveryDto struct {
	Email string `json:"email" validate:"required,email"`
}

func (c Controllers) HandlerPassRecovery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputPassRecoveryDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Fatalf("SI01: %s", err)
	}

	errors := c.InputValidation(input)

	if len(errors) > 0 {
		w.Write(errors)
		return
	}

	usecaseInput := usecase.InputPassRecoveryDto(input)

	var userRepoInterface repository.UserRepositoryInterface = &repository.UserRepository{}
	var userRecoveryPassInterface repository.UserRecoveryPassRepositoryInterface = &repository.UserRecoveryPassRepository{}
	var mailJobsRepoInterface repository.MailJobsRepositoryInterface = &repository.MailJobsRepository{}

	var useCase = usecase.PassRecoveryUsecase{}
	useCase.UserRepository = userRepoInterface
	useCase.UserRecoveryPassRepository = userRecoveryPassInterface
	useCase.MailJobsRepository = mailJobsRepoInterface

	output := useCase.PassRecovery(usecaseInput)

	var send outputControllerDto

	switch output.InternalStatus {
	case 0:
		send.Status = 0
		send.Message = "Email not found."
	case 2:
		send.Status = 0
		send.Message = "Internal Error"
	case 1:
		send.Status = 1
		send.Message = "Success"
	}

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Fatalf("SI02: %s", err)
	}

	w.Write(jsonResp)
}
