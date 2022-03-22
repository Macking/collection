package cmd

import (
	"net/http"
	"strconv"
)

func writeResponseJSON(w http.ResponseWriter, jsonBytes []byte, httpCode int) {
	w.Header().Set("Content-Length", strconv.Itoa(len(jsonBytes)))
	w.WriteHeader(httpCode)
	_, err := w.Write(jsonBytes)
	if err != nil {
		return
	}
}
