package controllers

import (
	"api-go/infra/jwt"
	validate "api-go/infra/validator"
	"encoding/json"
	"log"
	"net/http"
)

// import(
// signinController "api-user/controllers/signin"
// )

// type inputControllerDto struct{}

type outputControllerDto struct {
	Status  int64    `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
	Data    any      `json:"data,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

type Controllers struct{}

type ControllerInterface interface {
	HandlerHome(w http.ResponseWriter, r *http.Request)
	HandlerSignIn(w http.ResponseWriter, r *http.Request)
}

func authValidate(w http.ResponseWriter, r *http.Request) (uint64, bool) {
	var send outputControllerDto
	var jwtInterface jwt.JwtInterface = &jwt.Jwt{}
	userId, err := jwtInterface.VerifyJWT(w, r)

	if err != nil {
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())

		jsonResp, err := json.Marshal(send)

		if err != nil {
			log.Fatalf("Error in Json Marshal. %s", err)
		}

		w.Write(jsonResp)
		return 0, false
	}

	return userId, true
}

func (c *Controllers) InputValidation(w http.ResponseWriter, input any) bool {
	var send outputControllerDto
	var validator validate.ValidatorInterface = &validate.Validator{}
	errors := validator.InputValidator(input)

	if len(errors) > 0 {
		send.Status = 0
		send.Errors = errors

		resp, err := json.Marshal(send)

		if err != nil {
			log.Panic("SI02: %s", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(resp))
		return false
	}

	return true
}
