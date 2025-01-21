package wallet

import "github.com/gorilla/mux"

func SetWalletRouters(Router *mux.Router) {
	Router.HandleFunc("/{address}/balance", GetWalletInfoService).Methods("GET")
	Router.HandleFunc("/send", SendMoneyToWalletService).Methods("POST")
}
