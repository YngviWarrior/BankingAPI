package controllers

import (
	"encoding/json"
	repository "go-api/infra/database/repositories/mock"
	usecases "go-api/usecases/home"
	"log"
	"net/http"
)

type inputHomeDto struct{}

func (Controllers) HandlerHome(w http.ResponseWriter, r *http.Request) {
	var homeRepository repository.MockRepositoryInterface = repository.MockRepository{}
	var homeUseCase = usecases.HomeUseCase{}
	homeUseCase.HomeRepository = homeRepository

	output := homeUseCase.ListAll()

	var send outputControllerDto
	send.Status = 1
	send.Message = "Success"
	send.Data = output

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Fatalf("Error in Json Marshal. %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
