package transaction

import (
	"encoding/json"
	"errors"
	"infotecs_go/src/settings"
	"log"
	"net/http"
	"strconv"
)

func GetTransactionsInfoService(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	transactionsQuery := request.URL.Query().Get("count")
	if len(transactionsQuery) < 1 {
		settings.RaiseError(writer, request, "transactions number cannot be less than 0 or equal to 1", 400)
		return
	}
	transactionsNumber, err := strconv.Atoi(transactionsQuery)
	if err != nil {
		log.Println("Query parameter error: ", err)
		settings.QueryParamConvertError(writer, request)
		return
	}
	transactions, err := GetLastTransactionsRepository(transactionsNumber)
	if errors.Is(err, settings.TransactionNotFound) {
		settings.TransactionsDontFound(writer, request)
		return
	}
	if err != nil {
		log.Println("error while get transactions info: ", err)
		settings.RaiseError(writer, request, "get transactions info error", 400)
		return
	}
	err = json.NewEncoder(writer).Encode(transactions)
	if err != nil {
		log.Println("error while encoding transactions struct: ", err)
		settings.EncodingError(writer, request)
		return
	}
}
