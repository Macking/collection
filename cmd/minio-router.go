package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/klauspost/compress/gzhttp"
	"github.com/klauspost/compress/gzip"
	"net/http"
)

type minioAPIHandlers struct{}

func registerMinioRouter(router *mux.Router, enableConfigOps bool) {
	minioAPI := minioAPIHandlers{}
	// Admin router
	minioRouter := router.PathPrefix("/minio").Subrouter()
	gz, err := gzhttp.NewWrapper(gzhttp.MinSize(1000), gzhttp.CompressionLevel(gzip.BestSpeed))
	if err != nil {
		// Static params, so this is very unlikely.
		fmt.Println("Unable to initialize server: ", err)
	}
	minioRouter.Methods(http.MethodPost).Path("/check").
		HandlerFunc(gz(httpWrapper(minioAPI.MinioCheckMD5Handler))).
		Queries("md5", "{md5:.*}")
	minioRouter.Methods(http.MethodPost).Path("/upload").
		HandlerFunc(minioAPI.MinioUploadFileHandler).
		Queries("bucket", "{bucket:.*}").
		Queries("key", "{key:.*}").
		Queries("path", "{path:.*}")
}
