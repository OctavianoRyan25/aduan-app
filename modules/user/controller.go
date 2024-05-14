package user

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/base"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/helpers"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	useCase UseCase
}

func NewUserController(useCase UseCase) *UserController {
	return &UserController{
		useCase: useCase,
	}
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	// Bind the request body to the User struct
	user := new(User)
	ctx.Bind(user)
	user.Password = HashPass(user.Password)
	errCode, err := c.useCase.RegisterUser(user)
	response := MapToComplaintResponse(*user)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(errCode, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    response,
	}
	return ctx.JSON(constants.SuccessCode, successResponse)
}

func (c *UserController) LoginUser(ctx echo.Context) error {
	user := new(User)
	ctx.Bind(user)

	resp, errCode, err := c.useCase.LoginUser(user)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(errCode, errorResponse)
	}

	// fmt.Printf("Password yang diinputkan: %s\n", user.Password)
	// fmt.Printf("Password dari database: %s\n", resp.Password)

	comparePass := ComparePass([]byte(resp.Password), []byte(user.Password))
	if !comparePass {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeInvalidEmailorPassword,
			Message:   constants.ErrInvalidEmailorPassword,
		}
		return ctx.JSON(constants.ErrCodeInvalidEmailorPassword, errorResponse)
	}

	token, err := helpers.GenerateToken(uint(resp.ID), resp.Email, "user")
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}
	respToken := UserLoginResponse{
		Token: token,
	}

	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "User logged in successfully",
		Data:    respToken,
	}

	return ctx.JSON(constants.SuccessCode, successResponse)
}

func (c *UserController) InactiveUser(ctx echo.Context) error {
	// Get user_id from JWT token
	userID := ctx.Get("user_id").(uint)
	userIDInt := int(userID)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeFailParseID,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	errCode, err := c.useCase.InactiveUser(id, userIDInt)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(errCode, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "User inactive successfully",
	}
	return ctx.JSON(constants.SuccessCode, successResponse)
}
