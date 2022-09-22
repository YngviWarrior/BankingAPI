package api

import (
	"log"
	"net/http"
	"os"

	controllers "api-user/controllers"

	"github.com/gorilla/mux"
)

func InitServer() {
	r := mux.NewRouter()

	var controllersInterface controllers.ControllerInterface = controllers.Controllers{}

	r.HandleFunc("/public/home", controllersInterface.HandlerHome).Methods("GET")
	r.HandleFunc("/public/signin", controllersInterface.HandlerSignIn).Methods("GET")
	r.HandleFunc("/public/signup", controllersInterface.HandlerSignUp).Methods("POST")
	r.HandleFunc("/public/recoverpassword", controllersInterface.HandlerPassRecovery).Methods("POST")

	log.Printf("Running on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), r)

	if err != nil {
		log.Fatalf("%v", err)
	}
}
