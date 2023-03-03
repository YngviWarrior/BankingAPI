package api

import (
	"log"
	"net/http"
	"os"

	controllers "holder-ms/controllers"
	"holder-ms/infra/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitServer() {
	var d database.Database
	var database database.DatabaseInterface = &d
	database.CreatePool(150)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		Debug:            false,
	})

	r := mux.NewRouter()

	var controllersInterface controllers.ControllerInterface = &controllers.Controllers{DataBase: &d}

	r.HandleFunc("/holder/create", controllersInterface.HandlerCreateHolder).Methods("POST")
	r.HandleFunc("/holder/verify", controllersInterface.HandlerVerifyHolder).Methods("PUT")
	r.HandleFunc("/holder/delete", controllersInterface.HandlerDeleteHolder).Methods("DELETE")
	r.HandleFunc("/holder/find", controllersInterface.HandlerFindHolder).Methods("GET")

	log.Printf("Running on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), c.Handler(r))
	// err := http.ListenAndServe(os.Getenv("PORT"), r)

	if err != nil {
		log.Fatalf("%v", err)
	}

}
