package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	KeyId string `json:"keyId"`
}

func Hello(c echo.Context) error {

	headers := c.Request().Header
	var headersStr string

	// Loop through each header and format them
	for key, values := range headers {
		// Format each header key and value as "{key}  : {values} \n"
		for _, value := range values {
			headersStr += fmt.Sprintf("%s : %s\n", key, value)
		}
	}
	//xForwardedFor := c.Request().Header.Get("X-Forwarded-For")

	headersStr += fmt.Sprintf("X-Forwarded-For : %s\n", c.Request().RemoteAddr)

	// Return the formatted headers as a response
	return c.String(http.StatusOK, headersStr)
}
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

	var reqBody RequestBody

	// Bind the JSON request body to the struct
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	keyId := reqBody.KeyId

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
	// For testing
	// j, _ := DecryptAESGCM(k, encKey)
	// fmt.Println("j        ", string(j))

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
