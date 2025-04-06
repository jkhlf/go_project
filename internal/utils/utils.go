package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func ReadIDParam(r *http.Request) (int64, error) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		// This helps in debugging which URLs are causing issues
		return 0, fmt.Errorf("invalid id parameter (URL: %s)", r.URL.Path)
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id parameter format: %v", err)
	}

	return id, nil
}
