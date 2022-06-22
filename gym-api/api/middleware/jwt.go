package middleware

import (
	"gym-membership/api/common"
	"gym-membership/config"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	jwtSignedMethod = jwt.SigningMethodHS256
)

func JWTMiddleware(config *config.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.App.Key), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			var role string = claims["role"].(string)
			if ok && token.Valid {
				if role != "user" {
					return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
						Status:  "error",
						Code:    http.StatusForbidden,
						Message: "you dont have permission",
						Data:    nil,
					})
				}
			}

			exp := claims["exp"].(float64)
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "token expired",
					Data:    nil,
				})
			}

			return next(c)
		}
	}
}

func JWTMiddlewareOperator(config *config.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.App.Key), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			var role string = claims["role"].(string)
			if ok && token.Valid {
				if role != "operator" && role != "superadmin" {
					return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
						Status:  "error",
						Code:    http.StatusForbidden,
						Message: "you dont have permission",
						Data:    nil,
					})
				}
			}

			exp := claims["exp"].(float64)
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "token expired",
					Data:    nil,
				})
			}

			return next(c)
		}
	}
}

func JWTMiddlewareSuperAdmin(config *config.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.App.Key), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "invalid token",
					Data:    nil,
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			var role string = claims["role"].(string)
			if ok && token.Valid {
				if role != "superadmin" {
					return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
						Status:  "error",
						Code:    http.StatusForbidden,
						Message: "you dont have permission",
						Data:    nil,
					})
				}
			}

			exp := claims["exp"].(float64)
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Status:  "error",
					Code:    http.StatusForbidden,
					Message: "token expired",
					Data:    nil,
				})
			}

			return next(c)
		}
	}
}
