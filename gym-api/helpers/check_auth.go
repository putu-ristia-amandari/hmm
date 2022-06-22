package helpers

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckAuthorization(c echo.Context) (string, error) {
	signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(signature) < 2 {
		return "", errors.New("invalid token")
	}

	if signature[0] != "Bearer" {
		return "", errors.New("invalid token")
	}

	return signature[1], nil
}
