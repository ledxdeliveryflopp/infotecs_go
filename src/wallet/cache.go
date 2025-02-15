package wallet

import (
	"context"
	"fmt"
	"infotecs_go/src/settings"
	"log"
)

var (
	ctx = context.Background()
)

// SaveWalletInRedis saveWalletInRedis - Функция, сохранения кошелька в кэш Redis
//
// Аргументы: w *Wallet - указатель на структуру кошелька
func SaveWalletInRedis(w *Wallet) {
	key := fmt.Sprintf("wallet:%s", w.Number)
	err := settings.RedisClient().Set(ctx, key, w, 0).Err()
	if err != nil {
		log.Println("error while save wallet in redis")
		log.Printf("key: %s", key)
		log.Printf("error: %s", err)
	}
}

// GetWalletFromRedis - Функция, получение кошелька из кэша redis
//
// Аргументы: w *Wallet - указатель на структуру кошелька
//
// Возвращаемые значения - Сериализованную структуру Wallet или ошибку
func GetWalletFromRedis(walletNumber string) (Wallet, error) {
	var walletFromCache Wallet
	key := fmt.Sprintf("wallet:%s", walletNumber)
	err := settings.RedisClient().Get(ctx, key).Scan(&walletFromCache)
	if err != nil {
		log.Printf("error while get wallet from redis")
		log.Printf("key: %s", key)
		log.Printf("error: %s", err)
		return walletFromCache, err
	}
	return walletFromCache, err
}
