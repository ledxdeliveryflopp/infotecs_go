// Package transaction предоставляет функции для работы с транзакциями
package transaction

// Transaction - структура транзакции
type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
	Time      string `json:"time"`
}
