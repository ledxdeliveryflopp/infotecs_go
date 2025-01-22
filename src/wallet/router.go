// Package wallet предоставляет функции для работы с кошельками
package wallet

import "github.com/gorilla/mux"

// SetWalletRouters - Функция добавления эндпоинтов для работы с кошельками
//
// Аргументы: router *mux.Router - Основной роутер
func SetWalletRouters(Router *mux.Router) {
	Router.HandleFunc("/{address}/balance", GetWalletInfoService).Methods("GET")
	Router.HandleFunc("/send", SendMoneyToWalletService).Methods("POST")
}
