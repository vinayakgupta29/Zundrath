package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateKeyHandler(c echo.Context) error {
	auth := AuthoriseRequest(&c.Request().Header)
	if !auth {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	CreateKey()

	return c.JSON(http.StatusCreated, KeyArr)
}

func DeleteKeyHandler(c echo.Context) error {
	keyId := c.QueryParam("keyId")
	DeleteKey(KeyMetaData{KeyId: keyId})
	return c.JSON(http.StatusCreated, KeyArr)
}

func GetKeyHandler(c echo.Context) error {
	keyId := c.QueryParam("keyId")
	key, err := kProvider.GetKey(keyId)
	if len(keyId) <= 0 {
		return c.JSON(http.StatusPartialContent, map[string]string{"error": "Key Not given"})
	}
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusCreated, key)
}
func LogRequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the required information
		clientHeader := c.Request().Header.Get("Client-Header")
		authHeader := c.Request().Header.Get("Authorization")

		// Log the request
		c.Logger().Infof("%s %s %s", time.Now().Format(time.RFC3339), clientHeader, authHeader)

		return next(c)
	}
}
