package content

import (
	"gym-membership/api/common"
	"gym-membership/api/v1/content/response"
	contentBusiness "gym-membership/business/content"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	contentservice contentBusiness.Service
}

func CreateController(contentservice contentBusiness.Service) *Controller {
	return &Controller{
		contentservice: contentservice,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	contents, err := controller.contentservice.GetAllContent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetAllContentResponse(contents)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Get all contents successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	contentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	content, err := controller.contentservice.GetContentByID(contentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetContentByIDResponse(content)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "get content by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// func (controller *Controller) Create(c echo.Context) error {
// 	contents, err := controller.contentservice.CreateContent()
// 	if err := c.Bind(updateContentRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}
// }

// func (controller *Controller) Update(c echo.Context) error {
// 	contentID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	updateContentRequest := new(request.UpdateContentRequest)
// 	if err := c.Bind(updateContentRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	request := *updateContentRequest.ToSpec()

// 	content, err := controller.contentservice.UpdateContent(&request, contentID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusInternalServerError,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	response := response.CreateGetContentByIDResponse(content)
// 	responseData := common.DefaultDataResponse{
// 		Status:  "success",
// 		Code:    http.StatusOK,
// 		Message: "update content successfully",
// 		Data:    response,
// 	}

// 	return c.JSON(http.StatusOK, responseData)
// }

func (controller *Controller) Delete(c echo.Context) error {
	contentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	content, err := controller.contentservice.DeleteContent(contentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetContentByIDResponse(content)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "delete content successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
