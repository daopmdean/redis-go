package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("redis url schema:", os.Getenv("URL"))
	opt, err := redis.ParseURL(os.Getenv("URL"))
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	ctx := context.Background()

	err = client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("---val", val)

	session := map[string]string{
		"name":    "John",
		"surname": "Smith",
		"company": "Redis",
		"age":     "29",
	}
	for k, v := range session {
		err := client.HSet(ctx, "user-session:123", k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	userSession := client.HGetAll(ctx, "user-session:123").Val()
	fmt.Println("---userSession", userSession)
}
