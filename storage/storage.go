package storage

import "time"

type Service interface {
	Save(string, time.Time) error
	Get(string) (string, error)
	GetInfo(string) (*Item, error)
	Close() error
}

type Item struct {
	Id      uint64 `json:"id" redis:"id"`
	URL     string `json:"url" redis:"url"`
	AddedTime time.Time `json:"added_time" redis:"added_time"`
	ExpiresIn time.Duration `json:"expires" redis:"expires"`
	Visits  int    `json:"visits" redis:"visits"`
}

type ErrNotFound struct {}

func (e *ErrNotFound) Error() string {
	return "the key is not found"
}