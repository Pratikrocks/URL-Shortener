package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"urlShortener/config"
	"urlShortener/hadlers"
	"urlShortener/storage/redis"
	RD "github.com/go-redis/redis"
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

	redis.RedisDB = RD.NewClient(&RD.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if e != nil {
		panic(e)
	}
	rtr := hadlers.New()
	http.Handle("/", rtr)
	http.ListenAndServe(":8082", nil)
}
