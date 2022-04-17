package redis

import (
	json2 "encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"urlShortener/config"
	"urlShortener/storage"
)

type DB struct { Client *redis.Client }

func (db *DB) New(cfg config.Config) error {
	address := cfg.Redis.Addr
	password := cfg.Redis.Password
	Db := cfg.Redis.DB
	db.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       Db,
	})
	res, err := db.Client.Ping().Result()
	fmt.Println(res, err)
	return err
}

func (db *DB) Close() error {
	return db.Client.Close()
}

func (db * DB) Save(shortUrl string, metaData storage.Item) error {
	metaData.AddedTime = time.Now()
	json, err := json2.Marshal(metaData)
 	if err != nil {
		return err
	}
	err = db.Client.Set(shortUrl, json, time.Second * metaData.ExpiresIn).Err()
	return err
}

func (db * DB) Get(shortUrl string) (storage.Item, error) {
	json, err := db.Client.Get(shortUrl).Result()
	if err == redis.Nil {
		return storage.Item{}, &storage.ErrNotFound{}
	} else if err != nil {
		return storage.Item{}, err
	}
	var item storage.Item
	err = json2.Unmarshal([]byte(json), &item)
	if err != nil {
		return storage.Item{}, err
	}
	return item, nil
}

