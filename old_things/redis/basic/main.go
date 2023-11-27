package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

// Define a custom struct to hold Album data. Notice the struct tags?
// These indicate to redigo how to assign the data from the reply into
// the struct.
type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Fetch all album fields with the HGETALL command. Wrapping this
	// in the redis.Values() function transforms the response into type
	// []interface{}, which is the format we need to pass to
	// redis.ScanStruct() in the next step.
	values, err := redis.Values(conn.Do("HGETALL", "album:1"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(values))

	// Create an instance of an Album struct and use redis.ScanStruct()
	// to automatically unpack the data to the struct fields. This uses
	// the struct tags to determine which data is mapped to which
	// struct fields.
	var album Album
	err = redis.ScanStruct(values, &album)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", album)
}