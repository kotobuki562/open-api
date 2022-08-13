package main

import (
	"log"
	"net/http"
	"open-api/api"
)

type petHandler struct {}

func main() {
	http.Handle("/", api.Handler(&petHandler{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}