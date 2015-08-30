package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var (
	SqlDb *sql.DB
	WikiM *WikiwordModel
)

func Init() {

	log.Println("Delme: init")

	db, err := sql.Open("sqlite3", "wiki.db")
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	SqlDb = db

	SqlDb.SetMaxIdleConns(20)
	SqlDb.SetMaxOpenConns(30)

	WikiM = NewWikiModel()
	log.Println("Delme: NewWikiModel()", WikiM)
}
