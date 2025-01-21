package settings

import (
	"database/sql"
	"embed"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"log"
)

func ConnectToBD() *sql.DB {
	db, err := sql.Open("postgres", GetDatabaseUrl())
	if err != nil {
		panic(err)
	}
	return db
}

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
