package settings

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"log"
)

const (
	host     = "database"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "wallet"
)

func ConnectToBD() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}

func MigrateDatabase(migrationsPath embed.FS) {
	migrationsFolder := &migrate.EmbedFileSystemMigrationSource{FileSystem: migrationsPath, Root: "migrations"}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	migration, err := migrate.Exec(db, "postgres", migrationsFolder, migrate.Up)
	if err != nil {
		panic(err)
	}
	log.Printf("Apllied %d migration.", migration)
}
