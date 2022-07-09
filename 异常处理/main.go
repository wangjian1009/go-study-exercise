package main

import (
	"fmt"
)

const (
	dbUser = "root"
	dbPwd  = "wangjian1009"
)

func main() {
	db, err := openDb(dbUser, dbPwd)
	if err != nil {
		fmt.Printf("openDb error: %+v\n", err)
		return
	}
	defer db.Close()

	userMgr := NewUserManager(db)

	user, err := userMgr.FindUser(1234)
	if err != nil {
		fmt.Printf("find user fail: %+v\n", err)
		return
	}

	if user == nil {
		fmt.Printf("user not found")
	} else {
		fmt.Printf("user %d found, name is %s", user.Id, user.Name)
	}
}
