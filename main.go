package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var kProvider Key
var CONFIG = map[string]string{
	"KMS_STORE": "keystore",
}
var CLIENTID = map[string]string{
	"medoceua":  "xLxZj3QkcyER4X/MEkBE02b9Hdkgcm4ocekjKcpZcGk=",
	"hplus":     "XgNZYZVIgdjX0bK+EfS7PqIEQ3Zom3kp2kC5m80y1f8=",
	"doc-app":   "JyH6ZfMlghNjKez7FxRFfz4CQZBNeuQcxBmL2EJgIQQ=",
	"medocplus": "tBM9kW5lEo7aaoilj7eBZFmfOsZWWYTARshbCrI6MRc=",
}
var Mk MasterKey

func main() {
	Mk.MasterKey = Mk.GetMasterKey()

	os.Mkdir(CONFIG["KMS_STORE"], os.ModePerm)

	e := echo.New()

	os.Mkdir("logs", os.ModePerm)
	f, err := os.OpenFile("logs/applogs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${remote_ip} t=${latency_human} ${error} ${header:X-Client}\n",
	}))
	e.Use(middleware.Recover())

	e.Logger.SetOutput(f)
	e.GET("/", Hello)
	e.POST("/create", CreateKeyHandler)
	e.POST("/delete", DeleteKeyHandler)
	e.POST("/get", GetKeyHandler)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if _, ok := err.(*echo.HTTPError); ok {
			c.JSON(err.(*echo.HTTPError).Code, err.Error())
		}
	}
	e.Logger.Fatal(e.Start(":8080"))

}
