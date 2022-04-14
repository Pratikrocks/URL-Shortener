package redis

import (
	json2 "encoding/json"
	"github.com/go-redis/redis"
	"time"
	"urlShortener/config"
	"urlShortener/storage"
)

type DB struct { client *redis.Client }

func (db *DB) New(cfg config.Config) error {
	db.client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	_, err := db.client.Ping().Result()
	return err
}

func (db *DB) Close() error {
	return db.client.Close()
}

func (db * DB) Save(shortUrl string, metaData storage.Item) error {
	metaData.AddedTime = time.Now()
	json, err := json2.Marshal(metaData)
 	if err != nil {
		return err
	}
	err = db.client.Set(shortUrl, json, time.Second * metaData.ExpiresIn).Err()
	return err
}

