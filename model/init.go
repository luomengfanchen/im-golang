package model

import (
	"database/sql"
	"im/log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=go password=go dbname=im_db sslmode=disable")
	if err != nil {
		log.Error(err.Error())
	}
}