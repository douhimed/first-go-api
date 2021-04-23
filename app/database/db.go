package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)


type ProductDb interface {
	Open() error
	Close() error
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {

	mysql, err := sqlx.Open("mysql", mySqlConnStr)
	if err != nil {
		return err
	}

	log.Println("Connected to database ...")

	mysql.MustExec(createProductSchema)

	d.db = mysql
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}