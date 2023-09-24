package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "anhdao",
		"email": "anhdao@msu.edu",
	})
 
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email string
		FullName string
	}
	user := User{
		Email: "Anhdao@msu.edu",
		FullName: "Anhdao",
	}
	return c.JSON(http.StatusOK, user)
 
}