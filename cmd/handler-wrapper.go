package cmd

import "net/http"

func httpWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f.ServeHTTP(w, r)
		return
	}
}
