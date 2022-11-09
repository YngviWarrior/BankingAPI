package controllers

import (
	repository "api-go/infra/database/repositories/mock"
	usecases "api-go/usecases/home"
	"encoding/json"
	"log"
	"net/http"
)

func (*Controllers) HandlerHome(w http.ResponseWriter, r *http.Request) {
	var needAuth bool = true
	w.Header().Set("Content-Type", "application/json")

	_, ok := authValidate(w, r)
	if needAuth && !ok {
		return
	}

	var homeRepository repository.MockRepositoryInterface = repository.MockRepository{}
	var homeUseCase = usecases.HomeUseCase{}
	homeUseCase.HomeRepository = homeRepository

	output := homeUseCase.ListAll()

	var send outputControllerDto
	send.Status = 1
	send.Message = "Success"
	send.Errors = nil
	send.Data = output

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Fatalf("Error in Json Marshal. %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
