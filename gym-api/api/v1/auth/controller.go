package auth

import (
	"gym-membership/api/common"
	"gym-membership/api/v1/auth/request"
	"gym-membership/api/v1/auth/response"
	userBusiness "gym-membership/business/user"
	"gym-membership/config"
	"gym-membership/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	userService userBusiness.Service
	config      *config.AppConfig
}

func CreateController(userService userBusiness.Service, config *config.AppConfig) *Controller {
	return &Controller{
		userService: userService,
		config:      config,
	}
}

func (controller *Controller) Register(c echo.Context) error {
	var registerUserRequest *request.RegisterUserRequest
	if err := c.Bind(&registerUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	request := registerUserRequest.ToSpec()

	user, err := controller.userService.CreateUser(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateRegisterResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "register successfully, please check your email to verify your account",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

func (controller *Controller) Verify(c echo.Context) error {
	verifyCode := c.Param("token")
	if verifyCode == "" {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "verify code is empty",
			Data:    nil,
		})
	}

	user, err := controller.userService.VerifyEmail(verifyCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateVerifyResponse(user.Name, user.Email, user.VerifiedAt)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "verify successfully, please login",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Login(c echo.Context) error {
	var loginUserRequest *request.LoginUserRequest
	if err := c.Bind(&loginUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	request := loginUserRequest.ToSpec()

	user, err := controller.userService.GetUserByEmailAndPassword(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	claims := jwt.MapClaims{
		"userId": user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"role":   "user",
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	}

	JWTKey := []byte(controller.config.App.Key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(JWTKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateLoginResponse(user, jwt)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "login successfully",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

func (controller *Controller) GetProfile(c echo.Context) error {
	jwtToken, err := helpers.CheckAuthorization(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusForbidden,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	user, err := controller.userService.GetUserLogin(jwtToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusForbidden,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateProfileResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "get profile successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) UpdateProfile(c echo.Context) error {
	jwtToken, err := helpers.CheckAuthorization(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusForbidden,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	user, err := controller.userService.GetUserLogin(jwtToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusForbidden,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	updateUserRequest := new(request.UpdateProfileRequest)
	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	request := *updateUserRequest.ToSpec()

	user, err = controller.userService.UpdateUser(&request, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateProfileResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "update user successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) ChangePassword(c echo.Context) error {
	jwtToken, err := helpers.CheckAuthorization(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusForbidden,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	user, err := controller.userService.GetUserLogin(jwtToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusForbidden,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	updatePasswordRequest := new(request.UpdatePasswordRequest)
	if err := c.Bind(updatePasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	request := *updatePasswordRequest.ToSpec()

	user, err = controller.userService.UpdateUserPassword(&request, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateResetPasswordResponse(user.Name, user.Email)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Change password successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) ForgotPassword(c echo.Context) error {
	var forgotPasswordRequest *request.ForgotPasswordRequest

	if err := c.Bind(&forgotPasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	request := forgotPasswordRequest.ToSpec()

	user, err := controller.userService.ForgotPassword(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateForgotPasswordResponse(user.Name, user.Email)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "forgot password successfully, please check your email to verify your account",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) ResetPassword(c echo.Context) error {
	verifyCode := c.Param("token")
	if verifyCode == "" {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "verify code is empty",
			Data:    nil,
		})
	}

	var resetPasswordRequest *request.ResetPasswordRequest
	if err := c.Bind(&resetPasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	request := resetPasswordRequest.ToSpec()

	user, err := controller.userService.ResetPassword(request, verifyCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateResetPasswordResponse(user.Name, user.Email)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "reset password successfully, please login",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
