package controllers

import (
	"IskandarGo/helper"
	"IskandarGo/jwttoken"
	"IskandarGo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helper.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")

	password := c.FormValue("password")

	res, err := models.CheckLogin(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":   "500",
			"messages": "Email Or Password Invalid!!!",
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// token := jwt.New(jwt.SigningMethodHS384)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["email"] = email
	// claims["level"] = "admin"
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// t, err := token.SignedString([]byte("secret"))
	token, err := jwttoken.GenerateJWT(email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":   "500",
			"messages": err.Error(),
		})
	}
	id_admin, err := models.GetIdAdmin(email)
	println(id_admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":   "500",
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "200",
		"message": "Success",
		"token":   token,
		"isAdmin": "true",
	})
}
