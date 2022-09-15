package api

import (
	"log"
	"net/http"
	"os"

	homeControllers "go-api/controllers/home"
	signInControllers "go-api/controllers/signin"
	signUpcontrollers "go-api/controllers/signup"

	"github.com/gorilla/mux"
)

func InitServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeControllers.HandlerHome).Methods("GET")
	r.HandleFunc("/public/signin", signInControllers.HandlerSignIn).Methods("GET")
	r.HandleFunc("/public/signup", signUpcontrollers.HandlerSignUp).Methods("POST")

	log.Printf("Running on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), r)

	if err != nil {
		log.Fatalf("%v", err)
	}
}
