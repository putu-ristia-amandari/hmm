package user

import (
	"gym-membership/api/common"
	"gym-membership/api/v1/user/request"
	"gym-membership/api/v1/user/response"
	userBusiness "gym-membership/business/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service userBusiness.Service
}

func CreateController(service userBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	users, err := controller.service.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetAllUserResponse(users)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Get all users successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	user, err := controller.service.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetUserByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "get user by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	updateUserRequest := new(request.UpdateUserRequest)
	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	request := *updateUserRequest.ToSpec()

	user, err := controller.service.UpdateUser(&request, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetUserByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "update user successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) UpdatePassword(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
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

	user, err := controller.service.UpdatePassword(&request, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetUserByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "update user password successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Delete(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	user, err := controller.service.DeleteUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetUserByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "delete user successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
