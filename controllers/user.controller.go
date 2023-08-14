package controllers

import (
	"IskandarGo/helper"
	"IskandarGo/jwttoken"
	"IskandarGo/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func FectAllUsers(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	result, err := models.FetchAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func StoreUser(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	email := c.FormValue("email")
	phoneNumber := c.FormValue("phoneNumber")
	city := c.FormValue("city")
	zip := c.FormValue("zip")
	message := c.FormValue("message")
	password := c.FormValue("password")
	username := c.FormValue("username")
	address := c.FormValue("address")
	// password := helper.(password)
	hash, _ := helper.HashPassword(password)

	result, err := models.StoreUser(email, phoneNumber, city, zip, message, hash, username, address)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	userId := c.FormValue("userId")
	email := c.FormValue("email")
	phoneNumber := c.FormValue("phoneNumber")
	city := c.FormValue("city")
	zip := c.FormValue("zip")
	message := c.FormValue("message")
	password := c.FormValue("password")
	username := c.FormValue("username")
	address := c.FormValue("address")

	conv_id, err := strconv.Atoi(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(
		conv_id,
		email,
		phoneNumber,
		city,
		zip,
		message,
		password,
		username,
		address,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	userId := c.FormValue("userId")

	conv_id, err := strconv.Atoi(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(conv_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
