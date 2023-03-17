package main

import (
	"asso-api/config"
	"asso-api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ctx := config.Connexion()
	defer config.Client.Disconnect(ctx)
	r := mux.NewRouter()
	r.HandleFunc("/asso", handler.GetAssociations).Methods("GET").Queries("q", "{q}")

	http.ListenAndServe(":80", r)
}
