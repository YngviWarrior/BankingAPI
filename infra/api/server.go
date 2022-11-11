package api

import (
	"log"
	"net/http"
	"os"

	controllers "api-go/controllers"
	"api-go/infra/database"

	"github.com/gorilla/mux"
)

func InitServer() {
	var database database.DatabaseInterface = &database.Database{}
	database.CreatePool()

	r := mux.NewRouter()

	var controllersInterface controllers.ControllerInterface = &controllers.Controllers{}

	r.HandleFunc("/public/home", controllersInterface.HandlerHome).Methods("GET")
	r.HandleFunc("/public/signin", controllersInterface.HandlerSignIn).Methods("GET")

	log.Printf("Running on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), r)

	if err != nil {
		log.Fatalf("%v", err)
	}
}
