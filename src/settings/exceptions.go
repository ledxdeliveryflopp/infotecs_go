// Package settings предоставляет функции для настройки приложения
package settings

import (
	"log"
	"net/http"
)

// RaiseError - Функция, для вызова ошибки
//
// Аргументы: writer http.ResponseWriter, request *http.Request, detail string - информация о ошибке,
// code int - статус код ошибки
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func RaiseError(writer http.ResponseWriter, request *http.Request, detail string, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson(detail)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while error write json", err)
		return
	}
}

// WalletDontFound - Функция, для вызова ошибки которая сообщает что кошелька не существует
//
// Аргументы: writer http.ResponseWriter, request *http.Request,
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func WalletDontFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("wallet with this number don't found")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while error write json", err)
		return
	}
}

// TransactionsDontFound - Функция, для вызова ошибки которая сообщает что транзакции не найдены
//
// Аргументы: writer http.ResponseWriter, request *http.Request,
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func TransactionsDontFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("transactions don't found")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while error write json", err)
		return
	}
}

// EncodingError - Функция, для вызова ошибки которая сообщает что не удается десериализовать json
//
// Аргументы: writer http.ResponseWriter, request *http.Request,
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func EncodingError(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("error while encoding struct")
	if err != nil {
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

// QueryParamConvertError - Функция, для вызова ошибки которая сообщает что не удается конвертировать query параметр
//
// Аргументы: writer http.ResponseWriter, request *http.Request,
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func QueryParamConvertError(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("error while converting query param")
	if err != nil {
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

// NotEnoughMoneyInWallet - Функция, для вызова ошибки которая сообщает что недостаточно средств для перевода
//
// Аргументы: writer http.ResponseWriter, request *http.Request,
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func NotEnoughMoneyInWallet(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("wallet balance is less than requested")
	if err != nil {
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

// NotFoundEndpoint - Функция, для вызова ошибки которая сообщает что эндпоинт не найден
//
// Аргументы: writer http.ResponseWriter, request *http.Request,
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func NotFoundEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("endpoint not found")
	if err != nil {
		log.Println("error while build not found json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

// MethodNotAllowed - Функция, для вызова ошибки которая сообщает что эндпоинт не принимает такой метод запроса
//
// Аргументы: writer http.ResponseWriter, request *http.Request,
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или nil при ошибке сериализации
func MethodNotAllowed(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusMethodNotAllowed)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("method Not Allowed")
	if err != nil {
		log.Println("error while build not allowed json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while not allowed write json", err)
		return
	}
}
