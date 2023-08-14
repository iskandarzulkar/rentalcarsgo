package controllers

import (
	"IskandarGo/jwttoken"
	"IskandarGo/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func FectAllOrder(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	result, err := models.FetchAllOrder()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"pickUpTime": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func StoreOder(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	pickUpLoc := c.FormValue("pickUpLoc")
	dropOffLoc := c.FormValue("dropOffLoc")
	pickUpDate := c.FormValue("pickUpDate")
	dropOffDate := c.FormValue("dropOffDate")
	pickUpTime := c.FormValue("pickUpTime")
	carId := c.FormValue("carId")
	userId := c.FormValue("userId")
	adminId := c.FormValue("adminId")

	result, err := models.StoreOrder(pickUpLoc, dropOffLoc, pickUpDate, dropOffDate, pickUpTime, carId, userId, adminId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateOrder(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	orderId := c.Param("id")

	conv_id, err := strconv.Atoi(orderId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	pickUpLoc := keyVal["pickUpLoc"]
	dropOffLoc := keyVal["dropOffLoc"]
	pickUpDate := keyVal["pickUpDate"]
	dropOffDate := keyVal["dropOffDate"]
	pickUpTime := keyVal["pickUpTime"]
	carId := keyVal["carId"]
	userId := keyVal["userId"]
	adminId := keyVal["adminId"]

	println(pickUpTime)
	// _, err = stmt.Exec(first_name, last_name, email,
	// 	params["id"])
	// if err != nil {
	// 	panic(err.Error())
	// }

	// data := c.Request().Body

	// println(data)
	// pickUpLoc := c.FormValue("pickUpLoc")

	// dropOffLoc := c.FormValue("dropOffLoc")
	// // pickUpDate := c.Param("pickUpDate")
	// dropOffDate := c.FormValue("dropOffDate")
	// pickUpTime := c.FormValue("pickUpTime")
	// carId := c.FormValue("carId")
	// userId := c.FormValue("userId")
	// adminId := c.FormValue("adminId")

	result, err := models.UpdateOrder(
		conv_id,
		pickUpLoc,
		dropOffLoc,
		pickUpDate,
		dropOffDate,
		pickUpTime,
		carId,
		userId,
		adminId,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteOrder(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	orderId := c.Param("id")
	// orderId := c.FormValue("orderId")

	conv_id, err := strconv.Atoi(orderId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteOrder(conv_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func FetchCarsOrders(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	result, err := models.FetchCarsOrders()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"pickUpTime": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}
