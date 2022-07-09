package main

import (
	"database/sql"

	"github.com/pkg/errors"
)

func openDb(dbUser string, dbPwd string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dbUser+":"+dbPwd+"@tcp(127.0.0.1:3306)/beego?charset=utf8")
	if err != nil {
		return nil, errors.Wrap(err, "sql Open fail")
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, errors.Wrap(err, "sql link fail")
	}

	return
}
