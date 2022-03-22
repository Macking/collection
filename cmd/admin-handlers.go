package cmd

import (
	"context"
	"encoding/json"
	"github.com/Macking/collection/internal/madmin"
	"net/http"
	"strconv"
)

// ServerInfoHandler - GET /admin/info
func (a adminAPIHandlers) ServerInfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jsonBytes, err := json.Marshal(getServerInfo(ctx, r))
	if err != nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(""))
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(jsonBytes)))
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return
}

func getServerInfo(ctx context.Context, r *http.Request) madmin.InfoMessage {
	return madmin.InfoMessage{
		Message: "Hello",
	}
}
