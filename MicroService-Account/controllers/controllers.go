package controllers

import (
	"account-ms/infra/database"
	validate "account-ms/infra/validator"
	"encoding/json"
	"log"
	"net/http"
)

type outputControllerDto struct {
	Status  int64    `json:"status"`
	Message string   `json:"message,omitempty"`
	Data    any      `json:"data,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

type Controllers struct {
	DataBase *database.Database
}

type ControllerInterface interface {
	HandlerCreateAccount(w http.ResponseWriter, r *http.Request)
	HandlerFindAccount(w http.ResponseWriter, r *http.Request)
	HandlerDeleteAccount(w http.ResponseWriter, r *http.Request)
	HandlerBlockAccount(w http.ResponseWriter, r *http.Request)
	HandlerTransactionAccount(w http.ResponseWriter, r *http.Request)
	HandlerListStatement(w http.ResponseWriter, r *http.Request)
	HandlerListTransactionType(w http.ResponseWriter, r *http.Request)
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
			log.Panic("SI02: ", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(resp))
		return false
	}

	return true
}
