package models

import (
	"IskandarGo/db"
	"IskandarGo/helper"
	"database/sql"
	"fmt"
)

// type Admin struct {
// 	Id    int    `json:"id"`
// 	Email string `json:"email"`
// }

type (
	Admin struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}
)

func CheckLogin(email, password string) (bool, error) {
	var obj Admin
	var pwd string

	con := db.CreatCon()

	sqlStatement := "SELECT * FROM admin WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.Id, &obj.Email, &pwd,
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

func GetIdAdmin(email string) (isAdmin string, err error) {
	// println(email)
	// var obj Admin
	con := db.CreatCon()

	sqlStatement := "SELECT id FROM admin WHERE email = ?"

	rows, err := con.Query(sqlStatement, email)
	fmt.Println(rows)
	// sqlStatementCars := "select * from cars WHERE carId=?"

	// rows, err := con.Query(sqlStatementCars, carId)

	println(err)
	// expirationTime := time.Now().Add(72 * time.Hour)
	// claims := &JWTClaim{
	// 	Email: email,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err = token.SignedString(jwtKey)
	// ids := 7
	// isAdmin, err = res
	return
}
