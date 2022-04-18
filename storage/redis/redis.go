package redis

import (
	json2 "encoding/json"
	"github.com/go-redis/redis"
	"time"
	"urlShortener/storage"
)

var RedisDB *redis.Client

func Close() error {
	return RedisDB.Close()
}

func Save(shortUrl string, metaData storage.Item) error {
	metaData.AddedTime = time.Now()

	json, err := json2.Marshal(metaData)
 	if err != nil {
		return err
	}

	err = RedisDB.Set(shortUrl, json, 0).Err()
	if err != nil {
		return err
	}
	err = RedisDB.Set(metaData.URL, json, 0).Err()

	return err
}

func  Get(shortUrl string) (storage.Item, error) {
	json, err := RedisDB.Get(shortUrl).Result()
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

