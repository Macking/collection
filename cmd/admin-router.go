package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/klauspost/compress/gzhttp"
	"github.com/klauspost/compress/gzip"
	"net/http"
)

// adminAPIHandlers provides HTTP handlers for service admin API.
type adminAPIHandlers struct{}

func registerAdminRouter(router *mux.Router, enableConfigOps bool) {
	adminAPI := adminAPIHandlers{}
	// Admin router
	adminRouter := router.PathPrefix("/admin").Subrouter()
	gz, err := gzhttp.NewWrapper(gzhttp.MinSize(1000), gzhttp.CompressionLevel(gzip.BestSpeed))
	if err != nil {
		// Static params, so this is very unlikely.
		fmt.Println("Unable to initialize server: ", err)
	}
	adminRouter.Methods(http.MethodGet).Path("/info").HandlerFunc(gz(httpWrapper(adminAPI.ServerInfoHandler)))
}
