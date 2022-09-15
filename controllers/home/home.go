package homeControllers

import (
	"encoding/json"
	controllers "go-api/controllers"
	repository "go-api/infra/database/mock"
	usecases "go-api/usecases/home"
	"log"
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	var homeRepository repository.MockRepository
	var homeUseCase = usecases.HomeUseCase{}
	homeUseCase.HomeRepository = &homeRepository

	output := homeUseCase.ListAll()

	var send controllers.OutputControllerDto
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
