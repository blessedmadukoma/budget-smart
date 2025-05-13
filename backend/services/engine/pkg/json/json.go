package json

import (
	"encoding/json"
	"net/http"

	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return messages.ErrMissingRequestBody
	}

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return err
	}

	return nil
}

type ResponseMessage struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, message string, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	response := ResponseMessage{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}

	return json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, err.Error(), nil)
}
