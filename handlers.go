package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateKeyHandler(c echo.Context) error {
	auth := AuthoriseRequest(&c.Request().Header)
	if !auth {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	k, err := CreateKey()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, map[string]string{"keyId": k.KeyId})
}

func DeleteKeyHandler(c echo.Context) error {
	keyId := c.QueryParam("keyId")
	DeleteKey(KeyMetaData{KeyId: keyId})
	return c.JSON(http.StatusCreated, KeyArr)
}

func GetKeyHandler(c echo.Context) error {
	auth := AuthoriseRequest(&c.Request().Header)
	if !auth {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	keyId := c.QueryParam("keyId")
	if len(keyId) <= 0 {
		return c.JSON(http.StatusPartialContent, map[string]string{"error": "Key Not given"})
	}

	key, err := kProvider.GetKey(keyId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"error": err})
	}
	cheader := c.Request().Header.Get("X-Client")

	js, e0 := json.Marshal(key)
	if e0 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": e0})
	}
	encKey, e1 := base64.StdEncoding.DecodeString(CLIENTID[cheader])
	if e1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": e1})
	}
	k, er := EncryptAESGCM(encKey, js)
	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": er})
	}

	return c.JSON(http.StatusCreated, k)
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
