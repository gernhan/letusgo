package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type TestType struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

func main() {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	redisService := NewRedisService(pool, "0")
	ctx := context.Background()
	testObject := TestType{
		Name:  "Rong",
		Alias: "Long",
	}

	redisService.DeleteObjectWithKey("object1", ctx)
	// insert object
	redisService.SaveObjectWithKey(testObject, "object1", ctx)
	// retrieve object
	reflected, _ := redisService.GetObjectByKey("object1", ctx)
	cast := reflected.([]byte)
	var retrieved TestType
	_ = json.Unmarshal(cast, &retrieved)

	fmt.Println(retrieved.Name)
}
