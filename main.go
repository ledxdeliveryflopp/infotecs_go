package main

import (
	"embed"
	"github.com/gorilla/mux"
	"infotecs_go/src/settings"
	"infotecs_go/src/wallet"
	"log"
	"net/http"
)

//go:embed migrations/*.sql
var dbMigrations embed.FS

func main() {
	router := mux.NewRouter()
	ApiRouter := router.PathPrefix("/api").Subrouter()
	wallet.SetWalletRouters(ApiRouter)
	router.NotFoundHandler = http.HandlerFunc(settings.NotFoundEndpoint)
	router.MethodNotAllowedHandler = http.HandlerFunc(settings.MethodNotAllowed)
	settings.MigrateDatabase(dbMigrations)
	err := http.ListenAndServe(":1111", router)
	if err != nil {
		log.Println(err)
		return
	}
}
