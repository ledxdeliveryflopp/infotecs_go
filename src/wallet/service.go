// Package wallet предоставляет функции для работы с кошельками
package wallet

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"infotecs_go/src/settings"
	"io"
	"log"
	"net/http"
)

// GetWalletInfoService - Функция, для получения информации о кошельке
//
// Аргументы: writer http.ResponseWriter, request *http.Request
//
// Path параметры: address string - номер кошелька
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или Json с информацией о кошельке Wallet
func GetWalletInfoService(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	walletNumber := mux.Vars(request)["address"]
	if len(walletNumber) <= 1 {
		settings.RaiseError(writer, request, "wallet number cannot be less than 0 or equal to 1", 400)
		return
	}
	wallet, err := GetWalletByNumberRepository(walletNumber)
	switch {
	case errors.Is(sql.ErrNoRows, err):
		settings.WalletDontFound(writer, request)
		return
	case err != nil:
		log.Println("error while get wallet info: ", err)
		settings.RaiseError(writer, request, "get wallet info error", 400)
		return
	}
	err = json.NewEncoder(writer).Encode(wallet)
	if err != nil {
		log.Println("error while encoding wallet struct: ", err)
		settings.EncodingError(writer, request)
		return
	}
}

// SendMoneyToWalletService - Функция, для отправки денег на другой кошелек
//
// Аргументы: writer http.ResponseWriter, request *http.Request
//
// Тело запроса: from string - номер кошелька отправителя, to string - номер кошелька получателя
// amount float64 - сумма перевода
//
// Возвращаемые значения - Json с ошибкой settings.ErrorSchemas или Json с информацией BaseSchemas
func SendMoneyToWalletService(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var sendSchemas SendMoneySchemas
	err := sendSchemas.DecodeJson(request.Body)
	switch {
	case err == io.EOF:
		settings.RaiseError(writer, request, "empty request body", 400)
		return
	case err != nil:
		log.Println("error while decode wallet schemas: ", err)
		settings.EncodingError(writer, request)
		return
	}
	err = SendMoneyToWalletRepository(sendSchemas.From, sendSchemas.To, sendSchemas.Amount)
	switch {
	case errors.Is(settings.LowBalance, err):
		settings.NotEnoughMoneyInWallet(writer, request)
		return
	case errors.Is(sql.ErrNoRows, err):
		settings.WalletDontFound(writer, request)
		return
	case err != nil:
		log.Println("error while send money to wallet", err)
		settings.RaiseError(writer, request, "error while send money to wallet", 400)
		return
	}
	var response BaseSchemas
	marshaledResponse, err := response.BuildJson("success")
	if err != nil {
		log.Println("error while build response struct", err)
		settings.RaiseError(writer, request, "error while build response struct", 400)
		return
	}
	_, err = writer.Write(marshaledResponse)
	if err != nil {
		log.Println("error while write response struct", err)
		settings.RaiseError(writer, request, "write response error, but money transfer success", 400)
		return
	}
}
