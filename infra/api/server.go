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
	var d database.Database
	var database database.DatabaseInterface = &d
	database.CreatePool()

	r := mux.NewRouter()

	var controllersInterface controllers.ControllerInterface = &controllers.Controllers{DataBase: &d}

	r.HandleFunc("/public/home", controllersInterface.HandlerHome).Methods("GET")
	r.HandleFunc("/auth/signin", controllersInterface.HandlerSignIn).Methods("POST")

	log.Printf("Running on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), r)

	if err != nil {
		log.Fatalf("%v", err)
	}
}
