package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/riyanathariq/jwt-with-rbac/common"
	"github.com/riyanathariq/jwt-with-rbac/handler"
	"log"
	"net/http"
)

func Start() {
	r := mux.NewRouter()

	handler.InitHandler(r)

	fmt.Println(fmt.Sprintf("Starting server at port %s", common.GetEnv("APP_PORT", "8888")))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", common.GetEnv("APP_PORT", "8888")), r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
