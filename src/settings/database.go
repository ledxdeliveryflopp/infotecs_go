// Package settings предоставляет функции для настройки приложения
package settings

import (
	"database/sql"
	"embed"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"log"
)

// ConnectToBD - функция для подключения к бд
//
// Возвращаемые значения - *sql.DB при удачном подключении
//
// Вызывается panic() при ошибке подключения
func ConnectToBD() *sql.DB {
	db, err := sql.Open("postgres", GetDatabaseUrl())
	if err != nil {
		panic(err)
	}
	return db
}

// MigrateDatabase - функция для применения миграций
//
// # Аргументы - migrationsPath embed.FS - папка с миграциями
//
// Вызывается panic() при ошибке подключения к бд или ошибке применения миграций
func MigrateDatabase(migrationsPath embed.FS) {
	migrationsFolder := &migrate.EmbedFileSystemMigrationSource{FileSystem: migrationsPath, Root: "migrations"}
	log.Println(GetDatabaseUrl())
	db, err := sql.Open("postgres", GetDatabaseUrl())
	if err != nil {
		panic(err)
	}
	migration, err := migrate.Exec(db, "postgres", migrationsFolder, migrate.Up)
	if err != nil {
		panic(err)
	}
	log.Printf("Apllied %d migration.", migration)
}
