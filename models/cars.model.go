package models

import (
	"IskandarGo/db"
	"net/http"
)

type (
	Cars struct {
		CarId     int    `json:"carId"`
		Name      string `json:"name"  validate:"required"`
		CarType   string `json:"carType"  validate:"required"`
		Rating    string `json:"rating" validate:"required"`
		Fuel      string `json:"fuel" validate:"required"`
		Image     string `json:"image" validate:"required"`
		HourRate  string `json:"hourRate" validate:"required"`
		DayRate   string `json:"dayRate" validate:"required"`
		MonthRate string `json:"monthRate" validate:"required"`
	}

	// CustomValidator struct {
	// 	validator *validator.Validate
	// }
)

func FectAllCars() (Response, error) {
	var obj Cars

	var arrobj []Cars

	var res Response

	con := db.CreatCon()

	sqlStatement := "select * from cars"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&obj.CarId, &obj.Name, &obj.CarType, &obj.Rating, &obj.Fuel, &obj.Image, &obj.HourRate, &obj.DayRate, &obj.MonthRate)

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

func StoreCars(name string, carType string, rating string, fuel string, image string, hourRate string, dayRate string, monthRate string) (Response, error) {
	var res Response

	con := db.CreatCon()

	sqlStatement := "INSERT INTO cars (name, carType, rating, fuel, image, hourRate, dayRate, monthRate) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, carType, rating, fuel, image, hourRate, dayRate, monthRate)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	var obj Cars

	var arrobj []Cars

	sqlStatement2 := "select * from cars where carId=?"

	rows, err := con.Query(sqlStatement2, lastInsertedId)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&obj.CarId, &obj.Name, &obj.CarType, &obj.Rating, &obj.Fuel, &obj.Image, &obj.HourRate, &obj.DayRate, &obj.MonthRate)

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

func UpdateCars(carId int, name string, carType string, rating string, fuel string, image string, hourRate string, dayRate string, monthRate string) (Response, error) {

	var res Response
	con := db.CreatCon()

	sqlStatement := "UPDATE cars SET name= ?, carType= ?, rating= ?, fuel= ?, image= ?, hourRate= ?, dayRate= ?, monthRate= ? WHERE carId= ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(
		name,
		carType,
		rating,
		fuel,
		image,
		hourRate,
		dayRate,
		monthRate,
		carId,
	)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	println(rowsAffected)
	if err != nil {
		return res, err
	}

	var obj Cars

	var arrobj []Cars

	sqlStatementCars := "select * from cars WHERE carId=?"

	rows, err := con.Query(sqlStatementCars, carId)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&obj.CarId, &obj.Name, &obj.CarType, &obj.Rating, &obj.Fuel, &obj.Image, &obj.HourRate, &obj.DayRate, &obj.MonthRate)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	// res.Status = http.StatusOK
	// res.Message = "Success"

	// res.Data = map[string]int64{
	// 	"rows_affected": rowsAffected,
	// }

	return res, nil
}

func DeleteCars(carId int) (Response, error) {
	var res Response

	con := db.CreatCon()

	sqlStatement := "DELETE FROM `cars` WHERE carId = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(carId)

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
