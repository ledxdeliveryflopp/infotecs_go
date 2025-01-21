package settings

import (
	"embed"
	"fmt"
	"github.com/driftprogramming/godotenv"
	"log"
)

func InitSettings(envPath embed.FS) bool {
	err := godotenv.Load(envPath, "environment/.env")
	if err != nil {
		panic(err)
	}
	log.Println("Env inited")
	return true
}

func GetDatabaseUrl() string {
	host := godotenv.Get("HOST")
	port := godotenv.Get("PORT")
	user := godotenv.Get("USER")
	password := godotenv.Get("PASSWORD")
	dbname := godotenv.Get("NAME")
	databaseUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return databaseUrl
}
