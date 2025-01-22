// Package settings предоставляет функции для настройки приложения
package settings

import "errors"

// LowBalance - ошибка, которая сообщает что на кошельке недостаточно средств для перевода
var LowBalance = errors.New("wallet balance is less than requested")

// TransactionNotFound - ошибка, которая сообщает что не удается найти транзакции
var TransactionNotFound = errors.New("transactions not found")
