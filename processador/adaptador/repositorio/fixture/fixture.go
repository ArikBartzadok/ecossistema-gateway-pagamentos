package fixture

import (
	"context"
	"database/sql"
	"io/fs"
	"log"

	"github.com/maragudk/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func Subir(migrationsDiretorio fs.FS) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	if err := migrate.Up(context.Background(), db, migrationsDiretorio); err != nil {
		panic(err)
	}
	return db
}

func Descer(db *sql.DB, migrationsDiretorio fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDiretorio); err != nil {
		panic(err)
	}
	db.Close()
}
