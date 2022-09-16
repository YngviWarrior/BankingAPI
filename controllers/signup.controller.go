package controllers

import (
	"encoding/json"
	repository "go-api/infra/database/repositories/mysql"
	usecase "go-api/usecases/signup"
	"log"
	"net/http"
)

type inputSignUpDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Sponsor  string `json:"sponsor"`
}

func (c Controllers) HandlerSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputSignUpDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Fatalf("SI01: %s", err)
	}

	errors := c.InputValidation(input)

	if len(errors) > 0 {
		w.Write(errors)
		return
	}

	usecaseInput := usecase.InputSignUpDto(input)

	var repoInterface repository.UserRepositoryInterface = &repository.UserRepository{}
	var useCase = usecase.SignUpUsecase{}
	useCase.UserRepository = repoInterface

	output := useCase.SignUp(usecaseInput)

	var send outputControllerDto

	switch output.InternalStatus {
	case 0:
		send.Status = 0
		send.Message = "Email já cadastrado"
	case 2:
		send.Status = 0
		send.Message = "Failed to create"
	case 1:
		send.Status = 1
		send.Message = "Success"
		send.Data = output
	}

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Fatalf("SI02: %s", err)
	}

	w.Write(jsonResp)
}
