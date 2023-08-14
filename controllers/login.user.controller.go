package controllers

import (
	"IskandarGo/jwttoken"
	"IskandarGo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckLoginUser(c echo.Context) error {
	email := c.FormValue("email")

	password := c.FormValue("password")

	res, err := models.CheckLoginUser(email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":   "500",
			"messages": "Email Or Password Invalid!!!",
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token, err := jwttoken.GenerateJWT(email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":   "500",
			"messages": err.Error(),
		})
	}

	// data, err := models.GetUserByEmails(email)
	// fmt.Println(data[0])
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{
	// 		"status":   "500",
	// 		"messages": err.Error(),
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "200",
		"message": "Success",
		"token":   token,
		"isAdmin": "false",
	})

}
