package controllers

import (
	"encoding/json"
	repository "go-api/infra/database/repositories/mysql"
	usecase "go-api/usecases/signin"
	"log"
	"net/http"
)

type inputSignInDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (c Controllers) HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputSignInDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Fatalf("SI01: %s", err)
	}

	errors := c.InputValidation(input)

	if len(errors) > 0 {
		w.Write(errors)
		return
	}

	usecaseDto := usecase.InputSignInDto(input)

	var repoInterface repository.UserRepositoryInterface = &repository.UserRepository{}
	var useCase = usecase.SignInUsecase{}
	useCase.UserRepository = repoInterface

	output := useCase.SignIn(usecaseDto)

	var send outputControllerDto
	send.Status = 1
	send.Message = "Success"
	send.Data = output

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Fatalf("SI02: %s", err)
	}

	w.Write(jsonResp)
}
