package membership

import (
	"gym-membership/api/common"
	"gym-membership/api/v1/membership/response"
	memberBusiness "gym-membership/business/membership"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	memberservice memberBusiness.Service
}

func CreateController(memberservice memberBusiness.Service) *Controller {
	return &Controller{
		memberservice: memberservice,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	members, err := controller.memberservice.GetAllMember()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetAllMemberResponse(members)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Get all memberships successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	memberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	membership, err := controller.memberservice.GetMemberByID(memberID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetMembershipByIDResponse(membership)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "get membership by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// func (controller *Controller) Create(c echo.Context) error {
// 	members, err := controller.memberservice.CreateMember()
// 	if err := c.Bind(updateMemberRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}
// }

// func (controller *Controller) Update(c echo.Context) error {
// 	memberID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	updateMemberRequest := new(request.UpdateMemberRequest)
// 	if err := c.Bind(updateMemberRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusBadRequest,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	request := *updateMemberRequest.ToSpec()

// 	member, err := controller.memberservice.UpdateMember(&request, memberID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
// 			Status:  "error",
// 			Code:    http.StatusInternalServerError,
// 			Message: common.ErrorResponse(err),
// 			Data:    nil,
// 		})
// 	}

// 	response := response.CreateGetMembershipByIDResponse(member)
// 	responseData := common.DefaultDataResponse{
// 		Status:  "success",
// 		Code:    http.StatusOK,
// 		Message: "update membership successfully",
// 		Data:    response,
// 	}

// 	return c.JSON(http.StatusOK, responseData)
// }

func (controller *Controller) Delete(c echo.Context) error {
	memberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	user, err := controller.memberservice.DeleteMember(memberID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: common.ErrorResponse(err),
			Data:    nil,
		})
	}

	response := response.CreateGetMembershipByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "delete membership successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
