package common

import (
	"encoding/json"
	"net/http"
)

// SendJSON serializes the given object to JSON and sends it to the client.
func SendJSON(w http.ResponseWriter, e any, code ...int) {
	w.Header().Set("Content-Type", "application/json")
	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		SendError(w, err)
	}
}

func SendOK(w http.ResponseWriter, code ...int) {
	SendJSON(w, map[string]string{"message": "OK"}, code...)
}
