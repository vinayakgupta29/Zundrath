package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	name string `json:name`
	age  int    `json:age`
	tree string `json:tree`
}

func main() {
	CreateKey()
	e := echo.New()
	u := User{name: "John", age: 20, tree: ""}
	fmt.Print(u)
	e.GET("/create", func(c echo.Context) error {
		if err := c.Bind(&u); err != nil {
			return err
		}
		CreateKey()

		return c.JSON(http.StatusCreated, KeyArr)
	})
	e.GET("/delete", func(c echo.Context) error {
		keyId := c.QueryParam("keyId")
		DeleteKey(KeyMetaData{KeyId: keyId})
		return c.JSON(http.StatusCreated, KeyArr)
	})
	e.GET("/get", func(c echo.Context) error {
		keyId := c.QueryParam("keyId")
		key, err := GetKey(keyId)
		if len(keyId) <= 0 {
			return c.JSON(http.StatusPartialContent, "Key Not given")
		}
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusCreated, key)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
