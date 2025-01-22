package transaction

import (
	"database/sql"
	"fmt"
	"infotecs_go/src/settings"
	"log"
)

func CreateTransactionRepository(tx *sql.Tx, sender string, recipient string, amount float64) error {
	queryStr := fmt.Sprintf("INSERT INTO transaction (sender, recipient, amount) VALUES ('%s', '%s', %f)",
		sender, recipient, amount)
	_, err := tx.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

func GetLastTransactionsRepository(limit int) ([]Transaction, error) {
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
		return transactionSchemas, nil
	} else {
		return nil, settings.TransactionNotFound
	}
}
