package handler

import (
	"encoding/json"
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

type Response struct {
	Message string `json:"message"`
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Authenticated"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
