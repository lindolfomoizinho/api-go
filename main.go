package main

import (
	"net/http"

	"github.com/lindolfomoizinho/api-go/config"
	"github.com/lindolfomoizinho/api-go/handler"
)

func main() {

	datastore := config.NewDatastore()

	userH := handler.NewUserHandler(datastore)

	mux := http.NewServeMux()
	mux.Handle("/users", userH)
	mux.Handle("/users/", userH)

	http.ListenAndServe(":8080", mux)

}
