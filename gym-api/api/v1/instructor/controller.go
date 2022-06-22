package instructor

import (
	"gym-membership/api/common"
	"gym-membership/api/v1/instructor/response"
	instructorBusiness "gym-membership/business/instructor"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	instructorservice instructorBusiness.Service
}

func CreateController(instructorservice instructorBusiness.Service) *Controller {
	return &Controller{
		instructorservice: instructorservice,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	instructors, err := controller.instructorservice.GetAllInstructor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetAllInstructorResponse(instructors)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Get all instructors successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	instructorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	instructor, err := controller.instructorservice.GetInstructorByID(instructorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetInstructorByIDResponse(instructor)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "get instructor by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// func (controller *Controller) Create(c echo.Context) error {
// 	instructors, err := controller.instructorservice.CreateInstructor()
// 	if err := c.Bind(updateInstructorRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}
// }

// func (controller *Controller) Update(c echo.Context) error {
// 	instructorID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	updateInstructorRequest := new(request.UpdateInstructorRequest)
// 	if err := c.Bind(updateInstructorRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	request := *updateInstructorRequest.ToSpec()

// 	instructor, err := controller.instructorservice.UpdateInstructor(&request, instructorID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusInternalServerError,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	response := response.CreateGetInstructorByIDResponse(instructor)
// 	responseData := common.DefaultDataResponse{
// 		Status:  "success",
// 		Code:    http.StatusOK,
// 		Message: "update instructor successfully",
// 		Data:    response,
// 	}

// 	return c.JSON(http.StatusOK, responseData)
// }

func (controller *Controller) Delete(c echo.Context) error {
	instructorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	instructor, err := controller.instructorservice.DeleteInstructor(instructorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetInstructorByIDResponse(instructor)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "delete instructor successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
