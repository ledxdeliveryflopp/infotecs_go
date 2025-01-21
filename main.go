package main

import (
	"embed"
	"github.com/gorilla/mux"
	"infotecs_go/src/settings"
	"infotecs_go/src/transaction"
	"infotecs_go/src/wallet"
	"log"
	"net/http"
)

//go:embed environment/*
var env embed.FS

//go:embed migrations/*.sql
var dbMigrations embed.FS

func main() {
	router := mux.NewRouter()
	ApiRouter := router.PathPrefix("/api").Subrouter()
	wallet.SetWalletRouters(ApiRouter)
	transaction.SetTransactionRouters(ApiRouter)
	router.NotFoundHandler = http.HandlerFunc(settings.NotFoundEndpoint)
	router.MethodNotAllowedHandler = http.HandlerFunc(settings.MethodNotAllowed)
	settings.InitSettings(env)
	settings.MigrateDatabase(dbMigrations)
	err := http.ListenAndServe(":1111", router)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("application started")
}
