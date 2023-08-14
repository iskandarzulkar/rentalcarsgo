package controllers

import (
	"IskandarGo/jwttoken"
	"IskandarGo/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func FectAllCars(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	result, err := models.FectAllCars()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"image": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func ExStoreImage(c echo.Context) error {

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("Image Upload Err ->", err)
	}

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	fileExt := strings.Split(file.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", filename, fileExt)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("image/" + image)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, image)

}

func StoreCars(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	name := c.FormValue("name")
	carType := c.FormValue("carType")
	rating := c.FormValue("rating")
	fuel := c.FormValue("fuel")

	hourRate := c.FormValue("hourRate")
	dayRate := c.FormValue("dayRate")
	monthRate := c.FormValue("monthRate")

	// image := c.FormValue("image")

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("Image Upload Err ->", err)
	}

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	fileExt := strings.Split(file.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", filename, fileExt)
	imageUrl := fmt.Sprintf("http://localhost:8080/image/%s", image)
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("image/" + image)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	result, err := models.StoreCars(name, carType, rating, fuel, imageUrl, hourRate, dayRate, monthRate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

// func StoreCars(c echo.Context) error {
// 	name := c.FormValue("name")
// 	carType := c.FormValue("carType")
// 	rating := c.FormValue("rating")
// 	fuel := c.FormValue("fuel")
// 	image := c.FormValue("image")
// 	hourRate := c.FormValue("hourRate")
// 	dayRate := c.FormValue("dayRate")
// 	monthRate := c.FormValue("monthRate")

// 	result, err := models.StoreCars(name, carType, rating, fuel, image, hourRate, dayRate, monthRate)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

func UpdateCars(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	// carId := c.FormValue("carId")
	carId := c.Param("id")
	conv_id, err := strconv.Atoi(carId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	name := c.FormValue("name")
	carType := c.FormValue("carType")
	rating := c.FormValue("rating")
	fuel := c.FormValue("fuel")

	hourRate := c.FormValue("hourRate")
	dayRate := c.FormValue("dayRate")
	monthRate := c.FormValue("monthRate")
	file, err := c.FormFile("image")

	if err != nil {
		log.Println("Image Upload Err ->", err)
	}

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	fileExt := strings.Split(file.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", filename, fileExt)
	imageUrl := fmt.Sprintf("http://localhost:8080/image/%s", image)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("image/" + image)

	if err != nil {

		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// body, err := ioutil.ReadAll(c.Request().Body)

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// keyVal := make(map[string]string)
	// json.Unmarshal(body, &keyVal)

	// // name := keyVal["name"]
	// carType := keyVal["carType"]
	// rating := keyVal["rating"]
	// fuel := keyVal["fuel"]
	// image := keyVal["image"]
	// hourRate := keyVal["hourRate"]
	// dayRate := keyVal["dayRate"]
	// monthRate := keyVal["monthRate"]
	// println(image)

	// file, h, err := ioutil.ReadAll(c.Request().Body.FormFile("image"))
	// request.FormFile("photo")
	// println(file)
	// println(h)
	// uniqueId := uuid.New()

	// filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// fileExt := strings.Split(file.Filename, ".")[1]

	// image := fmt.Sprintf("%s.%s", filename, fileExt)

	// name := c.FormValue("name")
	// carType := c.FormValue("carType")
	// rating := c.FormValue("rating")
	// fuel := c.FormValue("fuel")
	// image := c.FormValue("image")
	// hourRate := c.FormValue("hourRate")
	// dayRate := c.FormValue("dayRate")
	// monthRate := c.FormValue("monthRate")

	// if err != nil {
	// 	log.Println("Image Upload Err ->", err)
	// }

	// uniqueId := uuid.New()

	// filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// fileExt := strings.Split(file.Filename, ".")[1]

	// image := fmt.Sprintf("%s.%s", filename, fileExt)
	// imageUrl := fmt.Sprintf("http://localhost:8080/image/%s", image)
	// src, err := file.Open()
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()

	// // Destination
	// dst, err := os.Create("image/" + image)
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// // Copy
	// if _, err = io.Copy(dst, src); err != nil {
	// 	return err
	// }

	// images, err := ioutil.ReadFile("images")
	// println(images)

	result, err := models.UpdateCars(conv_id, name, carType, rating, fuel, imageUrl, hourRate, dayRate, monthRate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCars(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.JSON(http.StatusInternalServerError, "Token Not Found")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	err := jwttoken.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteCars(conv_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":   "500",
			"messages": err.Error(),
		})

		// 	// return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
