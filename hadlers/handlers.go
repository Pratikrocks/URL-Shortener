package hadlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"urlShortener/base62"
	"urlShortener/storage"
	"urlShortener/storage/redis"
)

type Service struct {
	db *redis.DB
}

func (service *Service) New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping OK!!"))
	})
	mux.HandleFunc("/info", service.encode)
	return mux
}

func (service *Service)encode(w http.ResponseWriter,r *http.Request) {
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
	//service.db.Save(p.URL, p)
	w.Write([]byte(hash1))
}