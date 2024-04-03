package main

import "github.com/go-redis/redis/v8"

func IncRedis(redisClient *redis.Client) error {
	err := redisClient.Incr(ctx, "visitors").Err()
	if err != nil {
		return err
	}
	return nil
}
