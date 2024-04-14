package client

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestNewClientRedisClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:5001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Set(context.Background(), "key", "value", 0).Err(); err != nil {
		t.Fatal(err)
	}

	fmt.Println("WE ARE HERE")

	// val, err := rdb.Get(context.TODO(), "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(val)
}

func TestNewClient1(t *testing.T) {
	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	time.Sleep(time.Second)
	if err := c.Set(context.TODO(), "foo", 69); err != nil {
		log.Fatal(err)
	}
	val, err := c.Get(context.TODO(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	n, _ := strconv.Atoi(val)
	fmt.Println(n)
	fmt.Println("GET =>", val)
}

func TestNewClient(t *testing.T) {
	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("SET =>", fmt.Sprintf("bar_%d", i))
		if err := c.Set(context.TODO(), fmt.Sprintf("foo_%d", i), fmt.Sprintf("bar_%d", i)); err != nil {
			log.Fatal(err)
		}
		val, err := c.Get(context.TODO(), fmt.Sprintf("foo_%d", i))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("GET =>", val)
	}
}
