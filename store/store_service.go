package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"time"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 3 * time.Hour

func InitializeBackupStore() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "priyanshoon"
		password = "1212"
		dbname   = "url_shortner"
	)

	psqlconnect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconnect)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected successfully with database postgresql")
}

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("error init redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("failed saving key url | error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) (string, bool) {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		// panic(fmt.Sprintf("failed saving key url | error: %v - shortUrl: %s \n", err, shortUrl))
		return "", false
	}
	return result, true
}
