package transaction

import "github.com/gorilla/mux"

func SetTransactionRouters(Router *mux.Router) {
	Router.HandleFunc("/transactions", GetTransactionsInfoService).Methods("GET")
}
