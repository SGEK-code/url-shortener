package router

import "net/http"

func StartServer(addr string) error {
	mux := http.NewServeMux()
	addRoutes(mux)
	return http.ListenAndServe(addr, mux)
}
