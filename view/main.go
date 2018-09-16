package main

import (
	"os"
	
	"github.com/tariel-x/tsc/base"
	"github.com/go-redis/redis"
)

//go:generate tsc main.go DataIn DataOut

type DataIn struct {
	ID string `json:"id"`
}

type DataOut struct {
	Text string `json:"text"`
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
			text, err := client.Get(in.ID).Result()
			return DataOut{text}, err
		},
	)
	base.Die(err)
}