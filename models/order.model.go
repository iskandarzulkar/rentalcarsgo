package models

import (
	"IskandarGo/db"
	"net/http"
)

type (
	Order struct {
		OrderId     int    `json:"orderId"`
		PickUpLoc   string `json:"pickUpLoc"  validate:"required"`
		DropOffLoc  string `json:"dropOffLoc"  validate:"required"`
		PickUpDate  string `json:"pickUpDate" validate:"required"`
		DropOffDate string `json:"dropOffDate" validate:"required"`
		PickUpTime  string `json:"pickUpTime" validate:"required"`
		CarId       string `json:"carId" validate:"required"`
		UserId      string `json:"userId" validate:"required"`
		AdminId     string `json:"adminId" validate:"required"`
	}
	CarsOrders struct {
		OrderId     int    `json:"orderId"`
		PickUpLoc   string `json:"pickUpLoc"  validate:"required"`
		DropOffLoc  string `json:"dropOffLoc"  validate:"required"`
		PickUpDate  string `json:"pickUpDate" validate:"required"`
		DropOffDate string `json:"dropOffDate" validate:"required"`
		PickUpTime  string `json:"pickUpTime" validate:"required"`
		UserId      string `json:"userId" validate:"required"`
		AdminId     string `json:"adminId" validate:"required"`
		CarId       string `json:"carId"`
		Name        string `json:"name"  validate:"required"`
		CarType     string `json:"carType"  validate:"required"`
		Rating      string `json:"rating" validate:"required"`
		Fuel        string `json:"fuel" validate:"required"`
		Image       string `json:"image" validate:"required"`
		HourRate    string `json:"hourRate" validate:"required"`
		DayRate     string `json:"dayRate" validate:"required"`
		MonthRate   string `json:"monthRate" validate:"required"`
	}

	// CustomValidator struct {
	// 	validator *validator.Validate
	// }
)

func FetchAllOrder() (Response, error) {
	var obj Order

	var arrobj []Order

	var res Response

	con := db.CreatCon()

	sqlStatement := "select * from orders"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&obj.OrderId, &obj.PickUpLoc, &obj.DropOffLoc, &obj.PickUpDate, &obj.DropOffDate, &obj.CarId, &obj.UserId, &obj.AdminId, &obj.PickUpTime)

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

func StoreOrder(pickUpLoc string, dropOffLoc string, pickUpDate string, dropOffDate string, pickUpTime string, carId string, userId string, adminId string) (Response, error) {
	var res Response

	con := db.CreatCon()

	sqlStatement := "INSERT INTO orders (pickUpLoc, dropOffLoc, pickUpDate, dropOffDate, pickUpTime, userId, carId, adminId) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(pickUpLoc, dropOffLoc, pickUpDate, dropOffDate, pickUpTime, userId, carId, adminId)
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

func UpdateOrder(
	orderId int,
	pickUpLoc string,
	dropOffLoc string,
	pickUpDate string,
	dropOffDate string,
	pickUpTime string,
	carId string,
	userId string,
	adminId string,
) (Response, error) {
	var res Response
	con := db.CreatCon()

	sqlStatement := "UPDATE `orders` SET `pickUpLoc`=?,`dropOffLoc`=?,`pickUpDate`=?,`dropOffDate`=?,`pickUpTime`=?,`carId`=?,`userId`=?,`adminId`=? WHERE orderId=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(
		pickUpLoc,
		dropOffLoc,
		pickUpDate,
		dropOffDate,
		pickUpTime,
		carId,
		userId,
		adminId,
		orderId,
	)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	println(rowsAffected)

	if err != nil {
		return res, err
	}

	var obj CarsOrders

	var arrobj []CarsOrders

	sqlStatementOrders := "select * from qv_orders where orderId=?"

	rows, err := con.Query(sqlStatementOrders, orderId)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(
			&obj.OrderId,
			&obj.PickUpLoc,
			&obj.DropOffLoc,
			&obj.PickUpDate,
			&obj.DropOffDate,
			&obj.PickUpTime,
			&obj.UserId,
			&obj.AdminId,
			&obj.CarId,
			&obj.Name,
			&obj.CarType,
			&obj.Rating,
			&obj.Fuel,
			&obj.Image,
			&obj.HourRate,
			&obj.DayRate,
			&obj.MonthRate,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	// res.Data = map[string]int64{
	// 	"rows_affected": rowsAffected,
	// }

	return res, nil
}

func DeleteOrder(userId int) (Response, error) {
	var res Response

	con := db.CreatCon()

	sqlStatement := "DELETE FROM `orders` WHERE orderId = ?"

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

func FetchCarsOrders() (Response, error) {
	var obj CarsOrders

	var arrobj []CarsOrders

	var res Response

	con := db.CreatCon()

	sqlStatement := "select * from qv_orders"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(
			&obj.OrderId,
			&obj.PickUpLoc,
			&obj.DropOffLoc,
			&obj.PickUpDate,
			&obj.DropOffDate,
			&obj.PickUpTime,
			&obj.UserId,
			&obj.AdminId,
			&obj.CarId,
			&obj.Name,
			&obj.CarType,
			&obj.Rating,
			&obj.Fuel,
			&obj.Image,
			&obj.HourRate,
			&obj.DayRate,
			&obj.MonthRate,
		)

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
