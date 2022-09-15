package signinController

import (
	"encoding/json"
	controllers "go-api/controllers"
	repository "go-api/infra/database/mysql"
	usecases "go-api/usecases/signin"
	"log"
	"net/http"
)

func HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input usecases.InputSignInDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Fatalf("SI01: %s", err)
	}

	var repository repository.UserRepository
	var useCase = usecases.SignInUsecase{}
	useCase.UserRepository = &repository

	output := useCase.SignIn(input)

	var send controllers.OutputControllerDto
	send.Status = 1
	send.Message = "Success"
	send.Data = output

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Fatalf("SI02: %s", err)
	}

	w.Write(jsonResp)
}
