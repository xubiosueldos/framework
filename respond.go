package framework

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func RespondJSON(w http.ResponseWriter, status int, results interface{}) {

	response, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"codigo": strconv.Itoa(code), "error": message})
}
