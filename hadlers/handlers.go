package hadlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"urlShortener/base62"
	"urlShortener/storage"
	"urlShortener/storage/redis"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping OK!!"))
	})
	mux.HandleFunc("/info", encode)
	return mux
}

func encode(w http.ResponseWriter,r *http.Request) {
	fmt.Println("encode handlers: ", redis.RedisDB)
	var p storage.Item
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	p.AddedTime = time.Now()
	hash := rand.Intn(124567754)
	hash1 := base62.Encode(uint64(hash))
	json, err := json.Marshal(p)
	err = redis.RedisDB.Set(p.URL, json, 0).Err()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(hash1))
}