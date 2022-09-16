package controllers

import (
	"encoding/json"
	validate "go-api/infra/validator"
	"log"
	"net/http"
)

// import(
// signinController "go-api/controllers/signin"
// )

// type inputControllerDto struct{}

type outputControllerDto struct {
	Status  any    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

type Controllers struct{}

type ControllerInterface interface {
	HandlerHome(w http.ResponseWriter, r *http.Request)
	HandlerSignIn(w http.ResponseWriter, r *http.Request)
	HandlerSignUp(w http.ResponseWriter, r *http.Request)
	HandlerPassRecovery(w http.ResponseWriter, r *http.Request)
}

func (c Controllers) InputValidation(input any) (r []byte) {
	var send outputControllerDto
	var validator validate.ValidatorInterface = validate.Validator{}
	errors := validator.InputValidator(input)

	if len(errors) > 0 {
		send.Status = 0
		send.Message = "Error"
		send.Errors = errors

		resp, err := json.Marshal(send)

		if err != nil {
			log.Fatalf("SI02: %s", err)
		}

		return resp
	}

	return
}
