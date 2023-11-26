package store

import (
	"context"
	"fmt"
	"time"
	"github.com/go-redis/redis/v8"
)

type StorageService struct {
    redisClient *redis.Client
}

var (
    storeService = &StorageService{}
    ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
    redisClient := redis.NewClient(&redis.Options {
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    pong, err := redisClient.Ping(ctx).Result()
    if err != nil {
        panic(fmt.Sprintf("error init redis: %v", err))
    }

    fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
    storeService.redisClient = redisClient
    return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
    err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
    if err != nil {
        panic(fmt.Sprintf("failed saving key url | error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
    }
}

func RetrieveInitialUrl(shortUrl string) string {
    result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
    if err != nil {
        panic(fmt.Sprintf("failed saving key url | error: %v - shortUrl: %s \n", err, shortUrl))
    }

    return result
}
