package cmd

import (
	"github.com/gorilla/mux"
	"net/http"
)

func configureServerHandler() (http.Handler, error) {
	router := mux.NewRouter().SkipClean(true).UseEncodedPath()
	registerAdminRouter(router, true)
	registerMinioRouter(router, true)
	return router, nil
}
