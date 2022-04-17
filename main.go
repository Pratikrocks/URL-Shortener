package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"urlShortener/config"
	"urlShortener/hadlers"
	redis2 "urlShortener/storage/redis"
)

func main() {
	configJson, err := os.Open("configuration.json")
	if err != nil {
		panic(err)
	}
	defer configJson.Close()

	var cfg config.Config
	configBytes, e := io.ReadAll(configJson)
	if e != nil {
		panic(e)
	}
	json.Unmarshal(configBytes, &cfg)

	service := &hadlers.Service{}
	service.Db = &redis2.DB{}
	service.Db.New(cfg)

	if e != nil {
		panic(e)
	}
	mux := service.New()
	http.ListenAndServe(":8082", mux)
}
