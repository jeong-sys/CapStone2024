package main

import (
	"capstone.com/module/handler"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.POST("/signup", handler.SignUp)

}
