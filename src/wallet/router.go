package wallet

import "github.com/gorilla/mux"

func SetWalletRouters(Router *mux.Router) {
	Router.HandleFunc("/{address}/balance", GetWalletInfo).Methods("GET")
}
