package db

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
	"os"
	"pingo/models"
)

const databaseFilename = "./pingo.db"

type PingoDB struct {
	db gorp.DbMap
}

func (pdb *PingoDB) Init() error {

	_ = os.Remove(databaseFilename)

	db, err := sql.Open("sqlite3", databaseFilename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	pdb.db = gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	pdb.db.AddTable(models.Target{})

	return nil
}
