package controllers

import (
	repository "api-go/infra/database/repositories/mysql"
	usecase "api-go/usecases/signin"
	"encoding/json"
	"log"
	"net/http"
)

type inputSignInDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	IP       string
}

func (c *Controllers) HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input inputSignInDto
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

	input.IP = r.RemoteAddr
	usecaseInputDto := usecase.InputSignInDto(input)

	var repoInterface repository.RepositoriesInterface = &repository.Repositories{}
	var repoUserInterface repository.UserRepositoryInterface = &repository.UserRepository{}
	var useCase = usecase.SignInUsecase{}

	useCase.Repositories = repoInterface
	useCase.UserRepository = repoUserInterface

	output, err := useCase.SignIn(&usecaseInputDto)

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
