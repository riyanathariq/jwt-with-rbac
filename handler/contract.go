package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/riyanathariq/jwt-with-rbac/middleware"
	"net/http"
)

func InitHandler(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	api.Use(
		middleware.ProcessIdInjector,
	)

	api.HandleFunc("/test", TestHandler).
		Methods(http.MethodGet)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HIII!")
}
