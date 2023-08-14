package models

import (
	"IskandarGo/db"
	"IskandarGo/helper"
	"database/sql"
	"fmt"
)

type UsersLogin struct {
	UserId int    `json:"userId"`
	Email  string `json:"email"`
}

func CheckLoginUser(email, password string) (bool, error) {

	var obj UsersLogin
	var pwd string

	con := db.CreatCon()

	sqlStatement := "SELECT userId, email, password FROM users WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.UserId, &obj.Email, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("email Not Found!!!")
		return false, err
	}

	if err != nil {
		return false, err
	}

	match, err := helper.CheckPassword(password, pwd)

	if !match {
		fmt.Println("Password Dousen't Match")
		return false, err
	}

	return true, nil
}
