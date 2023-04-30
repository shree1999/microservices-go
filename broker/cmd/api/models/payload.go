package models

type JsonPayload struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}
