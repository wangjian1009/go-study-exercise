package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
	Id   int64
	Name string
}

type UserManager struct {
	db *sql.DB
}

func NewUserManager(db *sql.DB) (userMgr *UserManager) {
	return &UserManager{db: db}
}

func (um *UserManager) FindUser(id int64) (user *User, err error) {
	rows, err := um.db.Query("select id, name from tbl_user where id = id")
	if err != nil {
		return nil, errors.Wrap(err, "query user by id")
	}

	var found *User = nil

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				return nil, nil
			default:
				return nil, errors.Wrap(err, "query user by id")
			}
		}

		if found == nil {
			found = &user
		} else {
			return nil, errors.New(fmt.Sprintf("user %s duplicate", id))
		}
	}

	return found, nil
}
