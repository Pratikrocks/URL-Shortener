package hadlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
	"urlShortener/base62"
	"urlShortener/storage"
	"urlShortener/storage/redis"
)

func New() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping OK!!"))
	})
	rtr.HandleFunc("/info", encode)
	rtr.HandleFunc("/view/{id:[0-9a-zA-Z]+}", decode)
	return rtr
}

func encode(w http.ResponseWriter,r *http.Request) {
	fmt.Println("encode handlers: ", redis.RedisDB)

	var p storage.Item
	err := json.NewDecoder(r.Body).Decode(&p)

	check, _ := redis.Get(p.URL)

	if check.URL != "" {
		fmt.Println("URL already exists")
		w.Write([]byte("URL already exists"))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	p.AddedTime = time.Now()
	hash := rand.Intn(124567754)
	hash1 := base62.Encode(uint64(hash))

	for {
		if redis.RedisDB.Exists(hash1).Val() == 1 {
			item, err := redis.Get(hash1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			if item.URL == p.URL {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("URL already exists " + hash1))
				return
			}

			hash = rand.Intn(124567754)
			hash1 = base62.Encode(uint64(hash))
		} else {
			break
		}
	}

	err = redis.Save(hash1, p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(hash1))
}

func decode(w http.ResponseWriter, r *http.Request) {
	URL_Id := mux.Vars(r)
	urlId := URL_Id["id"]
	fmt.Println(urlId)

	item, err := redis.Get(urlId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, item.URL, 302)
}