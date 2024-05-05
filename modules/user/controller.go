package user

import (
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/base"
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
	err := c.useCase.RegisterUser(user)
	response := MapToComplaintResponse(*user)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
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

func (c *UserController) LoginUser(ctx echo.Context) error {
	user := new(User)
	ctx.Bind(user)

	resp, err := c.useCase.LoginUser(user)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return ctx.JSON(500, errorResponse)
	}

	// fmt.Printf("Password yang diinputkan: %s\n", user.Password)
	// fmt.Printf("Password dari database: %s\n", resp.Password)

	comparePass := ComparePass([]byte(resp.Password), []byte(user.Password))
	if !comparePass {
		errorResponse := base.ErrorResponse{
			Status:  "error",
			Message: "Password is incorrect",
		}
		return ctx.JSON(500, errorResponse)
	}

	token, err := GenerateToken(uint(resp.ID), resp.Email)
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
