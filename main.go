package main

import (
	"log"
	"net/http"
	"github.com/escherpad/api-server/services"
)

func main() {
	// There are two benefits for using a default db:
	// 1. database name need not be passed around or global or hardcoded
	// 2. avoid an extra var for db in trying to address the problems above
	services.DataStore.Init("localhost", "escherpad")
	services.DataStore.Session.DB("")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
