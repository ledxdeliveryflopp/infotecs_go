package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"infotecs_go/src/settings"
	"log"
)

var (
	ctx = context.Background()
)

// MarshalingTransactionsArray - Функция, сериализации массива транзакций
//
// Аргументы: t *[]Transaction - указатель на массив с транзакциями
//
// Возвращаемые значения - []byte при успешной сериализации или ошибку
func MarshalingTransactionsArray(t *[]Transaction) ([]byte, error) {
	result, err := json.Marshal(t)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UnmarshalTransactionArray - Функция, десериализации массива транзакций
//
// Аргументы: t *[]Transaction - указатель на массив с транзакциями(куда десериализовать), data string - строка из Redis
//
// Возвращаемые значения - error при ошибке десериализации
func UnmarshalTransactionArray(t *[]Transaction, data string) error {
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	return nil
}

// saveTransactionInRedis - Функция, сохранения массива транзакций в кэш Redis
//
// Аргументы: t *[]Transaction - указатель на массив с транзакциями, limit int - количество транзакций
func saveTransactionInRedis(t *[]Transaction, limit int) {
	result, err := MarshalingTransactionsArray(t)
	if err != nil {
		log.Println("error while marshaling transations list")
		log.Printf("error: %s", err)
	}
	key := fmt.Sprintf("transactions:%d", limit)
	err = settings.RedisClient().Set(ctx, key, result, 0).Err()
	if err != nil {
		log.Println("error while save transations list in redis")
		log.Printf("error: %s", err)
	}
}

// getTransactionFromRedis - Функция, получение массива с транзакциями из кэша Redis
//
// Аргументы: limit int - количество транзакций
//
// Возвращаемые значения - []Transaction при успешной десериализации информации из Redis или ошибку
func getTransactionFromRedis(limit int) ([]Transaction, error) {
	var transactions []Transaction
	key := fmt.Sprintf("transactions:%d", limit)
	result, err := settings.RedisClient().Get(ctx, key).Result()
	if err != nil {
		log.Println("error while get transactions from redis")
		log.Printf("key: %s", key)
		log.Printf("error: %s", err)
		return nil, err
	}
	err = UnmarshalTransactionArray(&transactions, result)
	if err != nil {
		log.Println("error while unmarshaling data from redis")
		log.Printf("data: %s", result)
		log.Printf("error: %s", err)
		return nil, err
	}

	return transactions, nil
}
