package signupController

import (
	"encoding/json"
	controllers "go-api/controllers"
	repository "go-api/infra/database/mysql"
	usecases "go-api/usecases/signup"
	"log"
	"net/http"
)

func HandlerSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input usecases.InputSignUpDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Fatalf("SI01: %s", err)
	}

	var repository repository.UserRepository
	var useCase = usecases.SignUpUsecase{}
	useCase.UserRepository = &repository

	output := useCase.SignUp(input)

	var send controllers.OutputControllerDto

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
