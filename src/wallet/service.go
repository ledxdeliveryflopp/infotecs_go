package wallet

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"infotecs_go/src/settings"
	"log"
	"net/http"
)

func GetWalletInfo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	walletNumber := mux.Vars(request)["address"]
	log.Println("address", walletNumber)
	if len(walletNumber) <= 1 {
		settings.RaiseError(writer, request, "wallet number cannot be less than 0 or equal to 1", 400)
		return
	}
	wallet, err := GetWalletByNumber(walletNumber)
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
