package main

import (
	"net/http"
	"urlShortener/hadlers"
)

func main() {
	var service hadlers.Service
	mux := service.New()
	http.ListenAndServe(":8082", mux)
}
