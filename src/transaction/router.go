// Package transaction предоставляет функции для работы с транзакциями
package transaction

import "github.com/gorilla/mux"

// SetTransactionRouters - Функция добавления эндпоинтов для работы с транзакциями
//
// Аргументы: router *mux.Router - Основной роутер
func SetTransactionRouters(router *mux.Router) {
	router.HandleFunc("/transactions", GetTransactionsInfoService).Methods("GET")
}
