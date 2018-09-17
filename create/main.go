package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"github.com/tariel-x/tsc/base"
)

//go:generate tsc main.go DataIn DataOut

type DataIn struct {
	NewText string `json:"newText"`
}

type DataOut struct {
	ID string `json:"id"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	base.Die(err)

	s, err := New(
		os.Getenv("RMQ"),
		os.Getenv("RMQ_API"),
		"create",
		"create",
		"view",
	)
	base.Die(err)

	err = s.Liftoff(
		func(in DataIn) (DataOut, error) {
			key := uuid.NewV4()
			err := client.Set(key.String(), in.NewText, 0).Err()
			return DataOut{key.String()}, err
		},
	)
	base.Die(err)
}
