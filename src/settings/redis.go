package settings

import "github.com/redis/go-redis/v9"

// RedisClient - Подключение к Redis
//
// Возвращаемые значения - *redis.Client указатель на клиент Redis
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "redis",
		DB:       0,
	})
	return client
}
