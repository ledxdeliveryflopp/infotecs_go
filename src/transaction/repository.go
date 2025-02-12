// Package transaction предоставляет функции для работы с транзакциями
package transaction

import (
	"database/sql"
	"fmt"
	"infotecs_go/src/settings"
	"log"
)

// CreateTransactionRepository - Функция, которая создает новый объект транзакции в бд
//
// Аргументы: tx *sql.Tx - транзакция бд, sender string - номер кошелька отправителя
// recipient string - номер кошелька получателя, amount float64 - сумма перевода
//
// Возвращаемые значения - error при ошибке добавления транзакции, nil при удачном создании транзакции
func CreateTransactionRepository(tx *sql.Tx, sender string, recipient string, amount float64) error {
	queryStr := fmt.Sprintf("INSERT INTO transaction (sender, recipient, amount) VALUES ('%s', '%s', %f)",
		sender, recipient, amount)
	_, err := tx.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

// GetLastTransactionsRepository - Функция, для получения N последних транзакций в бд
//
// Аргументы: limit int - количество выводимых транзакций
//
// Возвращаемые значения - error при ошибке получения транзакций, []Transaction при удачном получении транзакций
func GetLastTransactionsRepository(limit int) ([]Transaction, error) {
	TransactionsFromRedis, err := getTransactionFromRedis(limit)
	if err != nil {
		db := settings.ConnectToBD()
		queryStr := fmt.Sprintf("SELECT sender, recipient, amount, time FROM transaction ORDER BY time DESC LIMIT %d",
			limit)
		rows, err := db.Query(queryStr)
		if err != nil {
			log.Printf("query messages error: %s", err)
			return nil, err
		}
		defer rows.Close()
		var transactionSchemas []Transaction
		for rows.Next() {
			var t Transaction
			err := rows.Scan(&t.Sender, &t.Recipient, &t.Amount, &t.Time)
			if err != nil {
				log.Printf("rows scan error: %s", err)
				return nil, err
			}
			transactionSchemas = append(transactionSchemas, t)
		}
		if transactionSchemas != nil {
			saveTransactionInRedis(&transactionSchemas, limit)
			return transactionSchemas, nil
		} else {
			return nil, settings.TransactionNotFound
		}
	} else {
		return TransactionsFromRedis, nil
	}
}
