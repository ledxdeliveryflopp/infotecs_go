// Package settings предоставляет функции для настройки приложения
package settings

import (
	"embed"
	"fmt"
	"github.com/driftprogramming/godotenv"
	"log"
)

// InitSettings - Функция, для загрузки .env
//
// Аргументы: envPath embed.FS - путь к файлу настроек
//
// При ошибке вызывает panic()
func InitSettings(envPath embed.FS) {
	err := godotenv.Load(envPath, "environment/.env")
	if err != nil {
		panic(err)
	}
	log.Println("Env inited")
}

// GetDatabaseUrl - Функция, для создания url к бд
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
