package wallet

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"infotecs_go/src/settings"
	"io"
	"log"
	"net/http"
)

func GetWalletInfoService(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	walletNumber := mux.Vars(request)["address"]
	log.Println("address", walletNumber)
	if len(walletNumber) <= 1 {
		settings.RaiseError(writer, request, "wallet number cannot be less than 0 or equal to 1", 400)
		return
	}
	wallet, err := GetWalletByNumberRepository(walletNumber)
	if err != nil {
		log.Println("error while get wallet info", err)
		settings.RaiseError(writer, request, "get wallet info error", 400)
		return
	}
	err = json.NewEncoder(writer).Encode(wallet)
	if err != nil {
		log.Println("error while encoding wallet struct", err)
		settings.RaiseError(writer, request, "error while encoding wallet struct", 400)
		return
	}
}

func SendMoneyToWalletService(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var sendSchemas SendMoneySchemas
	decodedData, err := sendSchemas.decodeJson(request.Body)
	switch {
	case err == io.EOF:
		settings.RaiseError(writer, request, "empty request body", 400)
		return
	case err != nil:
		errorStr := fmt.Sprintf("error while decode: %s", err)
		settings.RaiseError(writer, request, errorStr, 400)
		return
	}
	err = SendMoneyToWalletRepository(decodedData.From, decodedData.To, decodedData.Amount)
	if err != nil {
		log.Println("error while send money to wallet", err)
		settings.RaiseError(writer, request, "error while send money to wallet", 400)
		return
	}
	var response BaseStruct
	builtResponse, err := response.buildJson("success")
	if err != nil {
		log.Println("error while build response struct", err)
		settings.RaiseError(writer, request, "error while build response struct", 400)
		return
	}
	_, err = writer.Write(builtResponse)
	if err != nil {
		log.Println("error while send response struct", err)
		settings.RaiseError(writer, request, "error while send response struct", 400)
		return
	}
}
