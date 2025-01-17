package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list..."))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created user"))
}

func main() {
	api := &api{addr: ":8080"}

	// Initialize a ServeMux
	mux := http.NewServeMux()

	srv := &http.Server{Addr: api.addr, Handler: mux}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	srv.ListenAndServe()
}
