package api

import (
	"gym-membership/api/middleware"
	authV1 "gym-membership/api/v1/auth"

	// memberV1 "gym-membership/api/v1/membership"
	userV1 "gym-membership/api/v1/user"
	"gym-membership/config"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller *userV1.Controller
	AuthV1Controller *authV1.Controller
	// MemberV1Controller *memberV1.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller, config *config.AppConfig) {
	userV1 := e.Group("/api/v1/user")
	userV1.GET("", controller.UserV1Controller.GetAll, middleware.JWTMiddlewareOperator(config))
	userV1.GET("/:id", controller.UserV1Controller.GetByID, middleware.JWTMiddlewareOperator(config))
	userV1.PUT("/:id", controller.UserV1Controller.Update, middleware.JWTMiddlewareOperator(config))
	userV1.PUT("/:id/password", controller.UserV1Controller.UpdatePassword, middleware.JWTMiddlewareOperator(config))
	userV1.DELETE("/:id", controller.UserV1Controller.Delete, middleware.JWTMiddlewareOperator(config))

	authV1 := e.Group("/api/v1/auth")
	authV1.POST("/register", controller.AuthV1Controller.Register)
	authV1.POST("/login", controller.AuthV1Controller.Login)
	authV1.GET("/profile", controller.AuthV1Controller.GetProfile, middleware.JWTMiddleware(config))
	authV1.PUT("/profile", controller.AuthV1Controller.UpdateProfile, middleware.JWTMiddleware(config))
	authV1.PUT("/profile/change-password", controller.AuthV1Controller.ChangePassword, middleware.JWTMiddleware(config))

	verifyV1 := e.Group("/api/v1/verify")
	verifyV1.GET("/:token", controller.AuthV1Controller.Verify)

	forgotPasswordV1 := e.Group("/api/v1/forgot-password")
	forgotPasswordV1.POST("", controller.AuthV1Controller.ForgotPassword)

	resetPasswordV1 := e.Group("/api/v1/reset-password")
	resetPasswordV1.POST("/:token", controller.AuthV1Controller.ResetPassword)
}

func MembershipPath(e *echo.Echo, controller Controller, config *config.AppConfig) {
	// memberV1 := e.Group("/api/v1/membership")
	// memberV1.GET("", controller.MemberV1Controller.GetAllMember)
	// memberV1.GET("/:id", controller.MemberV1Controller.GetMemberByID, middleware.JWTMiddleware(config))
	// memberV1.PUT("/:id", controller.MemberV1Controller.UpdateMember, middleware.JWTMiddleware(config))
	// memberV1.POST("/membership", controller.MemberV1Controller.CreateMember, middleware.JWTMiddleware(config))
}
