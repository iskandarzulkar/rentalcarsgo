package models

import (
	"IskandarGo/db"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// type Users struct {
// 	// *gorm.models
// 	UserId      int    `json:"userId"`
// 	Email       string `json:"email"  validate:"required"`
// 	PhoneNumber string `json:"phoneNumber"  validate:"required"`
// 	City        string `json:"city" validate:"required"`
// 	Zip         string `json:"zip" validate:"required"`
// 	Message     string `json:"message" validate:"required"`
// 	Password    string `json:"password" validate:"required"`
// 	Username    string `json:"username" validate:"required"`
// 	Address     string `json:"address" validate:"required"`
// }

type (
	Users struct {
		// *gorm.models
		UserId      int    `json:"userId"`
		Email       string `json:"email"  validate:"required"`
		PhoneNumber string `json:"phoneNumber"  validate:"required"`
		City        string `json:"city" validate:"required"`
		Zip         string `json:"zip" validate:"required"`
		Message     string `json:"message" validate:"required"`
		Password    string `json:"password" validate:"required"`
		Username    string `json:"username" validate:"required"`
		Address     string `json:"address" validate:"required"`
	}

	IdUsers struct {
		UserId int    `json:"userId"`
		Email  string `json:"email"  validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func FetchAllUsers() (Response, error) {
	var obj Users

	var arrobj []Users

	var res Response

	con := db.CreatCon()

	sqlStatement := "select * from users"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&obj.UserId, &obj.Email, &obj.PhoneNumber, &obj.City, &obj.Zip, &obj.Message, &obj.Password, &obj.Username, &obj.Address)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}
	fmt.Println(arrobj)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func GetUserByEmails2(email string) (Response, error) {

	var obj Users

	var arrobj []Users

	var res Response

	con := db.CreatCon()

	sqlStatement := "select * from users where email = ?"

	rows, err := con.Query(sqlStatement, email)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&obj.UserId, &obj.Email, &obj.PhoneNumber, &obj.City, &obj.Zip, &obj.Message, &obj.Password, &obj.Username, &obj.Address)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func GetUserByEmails(email string) ([]Users, error) {

	var obj Users

	var arrobj []Users

	con := db.CreatCon()

	sqlStatement := "select * from users where email = ?"
	rows, err := con.Query(sqlStatement, email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// var users []Users

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		// var obj Users

		if err = rows.Scan(
			&obj.UserId,
			&obj.Email,
			&obj.PhoneNumber,
			&obj.City,
			&obj.Zip,
			&obj.Message,
			&obj.Password,
			&obj.Username,
			&obj.Address); err != nil {
			return arrobj, err
		}
		// users = append(users, User{Id: id, Title: title})

		arrobj = append(arrobj, obj)
	}
	if err = rows.Err(); err != nil {
		return arrobj, err
	}

	return arrobj, nil
}

func StoreUser(email string, phoneNumber string, city string, zip string, message string, password string, username string, address string) (Response, error) {
	var res Response

	con := db.CreatCon()

	sqlStatement := "INSERT INTO users (email, phoneNumber, city, zip, message, password, username, address) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(email, phoneNumber, city, zip, message, password, username, address)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil

}

func UpdateUser(userId int, email string, phoneNumber string, city string, zip string, message string, password string, username string, address string) (Response, error) {
	var res Response
	con := db.CreatCon()

	sqlStatement := "UPDATE users SET email= ?, phoneNumber= ?, city= ?, zip= ?, message= ?,password= ?,username= ? ,address= ? WHERE userId= ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(
		email,
		phoneNumber,
		city,
		zip,
		message,
		username,
		password,
		address,
		userId,
	)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUser(userId int) (Response, error) {
	var res Response

	con := db.CreatCon()

	sqlStatement := "DELETE FROM `users` WHERE userId = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(userId)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}
