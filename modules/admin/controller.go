package admin

import (
	"strconv"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/base"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/helpers"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"github.com/labstack/echo/v4"
)

type AdminController struct {
	useCase UseCase
}

func NewAdminController(useCase UseCase) *AdminController {
	return &AdminController{
		useCase: useCase,
	}
}

func (c *AdminController) RegisterAdmin(ctx echo.Context) error {
	// Bind the request body to the Admin struct
	admin := new(Admin)
	ctx.Bind(admin)
	admin.Password = HashPass(admin.Password)
	errCode, err := c.useCase.RegisterAdmin(admin)
	response := MapToComplaintResponse(*admin)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    response,
	}
	return ctx.JSON(200, successResponse)
}

func (c *AdminController) LoginAdmin(ctx echo.Context) error {
	user := new(Admin)
	ctx.Bind(user)

	resp, errCode, err := c.useCase.LoginAdmin(user)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}

	// fmt.Printf("Password yang diinputkan: %s\n", user.Password)
	// fmt.Printf("Password dari database: %s\n", resp.Password)

	comparePass := ComparePass([]byte(resp.Password), []byte(user.Password))
	if !comparePass {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeInvalidEmailorPassword,
			Message:   "Password is incorrect",
		}
		return ctx.JSON(500, errorResponse)
	}

	token, err := helpers.GenerateToken(uint(resp.ID), resp.Email, "admin")
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}

	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "User logged in successfully",
		Data:    token,
	}

	return ctx.JSON(200, successResponse)
}

func (c *AdminController) UpdateStatusComplaint(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeFailParseID,
			Message:   err.Error(),
		}
		return ctx.JSON(400, errorResponse)
	}
	statusID := ctx.FormValue("status_id")
	conv, err := strconv.Atoi(statusID)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeFailParseID,
			Message:   err.Error(),
		}
		return ctx.JSON(400, errorResponse)

	}
	errCode, err := c.useCase.UpdateStatusComplaint(id, conv)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Complaint status updated successfully",
	}
	return ctx.JSON(200, successResponse)
}

func (c *AdminController) GetAllComplaint(ctx echo.Context) error {
	role := ctx.Get("role").(string)
	if role != "admin" {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeUnauthorized,
			Message:   constants.ErrUnauthorized,
		}
		return ctx.JSON(401, errorResponse)
	}
	complaints, errCode, err := c.useCase.GetAllComplaint()
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Complaints retrieved successfully",
		Data:    complaints,
	}
	return ctx.JSON(200, successResponse)
}

func (c *AdminController) GetAllUser(ctx echo.Context) error {
	role := ctx.Get("role").(string)
	if role != "admin" {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeUnauthorized,
			Message:   constants.ErrUnauthorized,
		}
		return ctx.JSON(401, errorResponse)
	}
	users, errCode, err := c.useCase.GetAllUser()
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    users,
	}
	return ctx.JSON(200, successResponse)
}

func (c *AdminController) UpdatePasswordUser(ctx echo.Context) error {
	role := ctx.Get("role").(string)
	if role != "admin" {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeUnauthorized,
			Message:   constants.ErrUnauthorized,
		}
		return ctx.JSON(401, errorResponse)
	}
	user := new(user.User)
	ctx.Bind(user)
	user.Password = HashPass(user.Password)
	errCode, err := c.useCase.UpdatePasswordUser(user)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Password updated successfully",
	}
	return ctx.JSON(200, successResponse)
}

func (c *AdminController) ActivateUser(ctx echo.Context) error {
	role := ctx.Get("role").(string)
	if role != "admin" {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeUnauthorized,
			Message:   constants.ErrUnauthorized,
		}
		return ctx.JSON(401, errorResponse)
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeFailParseID,
			Message:   err.Error(),
		}
		return ctx.JSON(400, errorResponse)
	}
	_, err = c.useCase.ActivateUser(id)
	if err != nil {
		return ctx.JSON(500, base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrorCodeBadRequest,
			Message:   err.Error(),
		})
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "User Activated",
	}
	return ctx.JSON(200, successResponse)
}
