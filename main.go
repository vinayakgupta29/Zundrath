package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Tree string `json:"tree"`
}

var kProvider Key

func main() {
	iv := make([]byte, 16)
	rand.Read(iv)
	fmt.Println(iv)
	e := echo.New()

	f, err := os.OpenFile("applogs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency} ${error}\n",
	}))
	e.Use(middleware.Recover())

	e.Logger.SetOutput(f)
	e.GET("/create", CreateKeyHandler)
	e.GET("/delete", DeleteKeyHandler)
	e.GET("/get", GetKeyHandler)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if _, ok := err.(*echo.HTTPError); ok {
			c.JSON(err.(*echo.HTTPError).Code, err.Error())
		}
	}
	e.Logger.Fatal(e.Start(":8080"))
}
