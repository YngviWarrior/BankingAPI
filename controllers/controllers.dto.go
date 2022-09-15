package controllers

type InputControllerDto struct{}

type OutputControllerDto struct {
	Status  any    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
