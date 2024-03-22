package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, v interface{}) error {
	fmt.Println("Parsing JSON", r.Body)
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(v)
	if err != nil {
		return
	}
}
