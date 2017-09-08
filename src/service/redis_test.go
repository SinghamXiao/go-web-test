package service

import (
	"testing"
	"fmt"
)

func TestRedisCache(t *testing.T) {
	redisClient, err := NewRedisClient(`{"conn": "127.0.0.1:6379"}`)
	if err != nil {
		t.Error("Init err")
	}

	err = redisClient.Set("key1", 1)
	if err != nil {
		t.Error("Set err")
	}

	err = redisClient.Set("key2", "2")
	if err != nil {
		t.Error("Set err")
	}

	v, err := redisClient.Get("key1")
	if err != nil {
		t.Error("Set err")
	}
	fmt.Println("key1: ", string(v))

	v, err = redisClient.Get("key2")
	if err != nil {
		t.Error("Set err")
	}
	fmt.Println("key2: ", string(v))

}