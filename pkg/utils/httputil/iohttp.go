package httputil

import (
	"encoding/json"
	"net/http"
)

func SendAnswer(data interface{}, w http.ResponseWriter) error {
	if resp, err := json.Marshal(data); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	} else {
		return err

	}
	return nil
}