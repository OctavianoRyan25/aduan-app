package complaint

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/base"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/api"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/storage"
	"github.com/labstack/echo/v4"
)

type ComplaintController struct {
	UseCase UseCase
	storage *storage.Storage
}

func NewComplaintController(UseCase UseCase, Storage *storage.Storage) *ComplaintController {
	return &ComplaintController{
		UseCase: UseCase,
		storage: Storage,
	}
}

func (c *ComplaintController) CreateComplaint(ctx echo.Context) error {
	// Get user_id from JWT token
	userID := ctx.Get("user_id").(uint)
	userIDInt := int(userID)
	//Get user role from JWT token
	role := ctx.Get("role").(string)
	if role != "user" {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeUnauthorized,
			Message:   constants.ErrUnauthorized,
		}
		return ctx.JSON(401, errorResponse)
	}
	// Get form value
	name := ctx.FormValue("name")
	phone := ctx.FormValue("phone")
	body := ctx.FormValue("body")
	category := ctx.FormValue("category")
	latitude := ctx.FormValue("latitude")
	longitude := ctx.FormValue("longitude")

	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrorCodeBadRequest,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	long, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrorCodeBadRequest,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	// Reverse geocoding
	location, err := api.ReverseGeocode(lat, long)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status: "error",
			// ErrorCode: constants.ErrorCodeBadRequest,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)

	}
	// Parse multipart form for images
	form, err := ctx.MultipartForm()
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrorCodeBadRequest,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	files := form.File["images"]
	var imagePaths []string

	// Save images
	for _, img := range files {
		ext := filepath.Ext(img.Filename)
		imagePath := "public/" + generateUniqueFilename(ext)

		// Open the uploaded file
		file, err := img.Open()
		if err != nil {
			errorResponse := base.ErrorResponse{
				Status:    "error",
				ErrorCode: constants.ErrorCodeBadRequest,
				Message:   err.Error(),
			}
			return ctx.JSON(http.StatusInternalServerError, errorResponse)
		}
		defer file.Close()

		// Save image data to storage
		if err := c.storage.SaveImage(imagePath, file); err != nil {
			errorResponse := base.ErrorResponse{
				Status:    "error",
				ErrorCode: constants.ErrorCodeBadRequest,
				Message:   err.Error(),
			}
			return ctx.JSON(http.StatusInternalServerError, errorResponse)
		}

		// Append image path to the list
		imagePaths = append(imagePaths, imagePath)
	}

	// Create Complaint struct
	complaint := &Complaint{
		Name:     name,
		Phone:    phone,
		Body:     body,
		Category: category,
		Images:   make([]Image, len(imagePaths)),
		UserID:   userIDInt,
		Location: location,
	}

	// Set image paths
	for i, path := range imagePaths {
		complaint.Images[i] = Image{Path: path}
	}

	// Call use case to create complaint
	errCode, err := c.UseCase.CreateComplaint(complaint)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := mapToCreateComplaintResponse(*complaint)

	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Complaint created successfully",
		Data:    response,
	}

	return ctx.JSON(http.StatusCreated, successResponse)
}

func (c *ComplaintController) GetAllComplaint(ctx echo.Context) error {
	// Get user_id from JWT token
	userID := ctx.Get("user_id").(uint)
	userIDInt := int(userID)
	//Get user role from JWT token
	role := ctx.Get("role").(string)
	if role != "user" {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrCodeUnauthorized,
			Message:   constants.ErrUnauthorized,
		}
		return ctx.JSON(401, errorResponse)
	}
	complaints, errCode, err := c.UseCase.GetAllComplaint(userIDInt)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	var complaintResponses []ComplaintResponse
	for _, complaint := range complaints {
		complaintResponses = append(complaintResponses, mapToComplaintResponse(complaint))
	}
	//Map complaints to response
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Success get complaint",
		Data:    complaintResponses,
	}
	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *ComplaintController) GetComplaintByID(ctx echo.Context) error {
	// Get user_id from JWT token
	userID := ctx.Get("user_id").(uint)
	userIDInt := int(userID)
	//Get user role from JWT token
	role := ctx.Get("role").(string)
	if role != "user" {
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
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}
	complaint, errCode, err := c.UseCase.GetComplaintByID(id, userIDInt)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	complaintResponse := mapToComplaintResponse(*complaint)

	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Success get complaint by id",
		Data:    complaintResponse,
	}
	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *ComplaintController) UpdateComplaint(ctx echo.Context) error {
	// Get user_id from JWT token
	userID := ctx.Get("user_id").(uint)
	userIDInt := int(userID)
	//Get user role from JWT token
	role := ctx.Get("role").(string)
	if role != "user" {
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
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	// Get form value
	name := ctx.FormValue("name")
	phone := ctx.FormValue("phone")
	body := ctx.FormValue("body")
	category := ctx.FormValue("category")
	latitude := ctx.FormValue("latitude")
	longitude := ctx.FormValue("longitude")

	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrorCodeBadRequest,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	long, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrorCodeBadRequest,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	// Reverse geocoding
	location, err := api.ReverseGeocode(lat, long)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: constants.ErrorCodeBadRequest,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	//complaint := new(Complaint)
	// if err := ctx.Bind(complaint); err != nil {
	// 	errorResponse := base.ErrorResponse{
	// 		Status:  "error",
	// 		Message: err.Error(),
	// 	}
	// 	return ctx.JSON(http.StatusBadRequest, errorResponse)
	// }

	// Create Complaint struct
	complaint := &Complaint{
		Name:     name,
		Phone:    phone,
		Body:     body,
		Category: category,
		Location: location,
	}

	errCode, err := c.UseCase.UpdateComplaint(id, userIDInt, complaint)

	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Complaint updated successfully",
	}
	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *ComplaintController) DeleteComplaint(ctx echo.Context) error {
	// Get user_id from JWT token
	userID := ctx.Get("user_id").(uint)
	userIDInt := int(userID)
	//Get user role from JWT token
	role := ctx.Get("role").(string)
	if role != "user" {
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
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	errCode, err := c.UseCase.DeleteComplaint(id, userIDInt)
	if err != nil {
		errorResponse := base.ErrorResponse{
			Status:    "error",
			ErrorCode: errCode,
			Message:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	successResponse := base.SuccessResponse{
		Status:  "success",
		Message: "Complaint deleted successfully",
	}
	return ctx.JSON(http.StatusOK, successResponse)
}
