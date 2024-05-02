package complaint

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

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
	// Bind the request body to the Complaint struct
	complaint := new(Complaint)
	complaint.Name = ctx.FormValue("name")
	complaint.Phone = ctx.FormValue("phone")
	complaint.Body = ctx.FormValue("body")
	complaint.Category = ctx.FormValue("category")
	form, err := ctx.MultipartForm()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	files := form.File["images"]

	// Save images
	for _, img := range files {
		// Generate unique filename
		ext := filepath.Ext(img.Filename)
		imagePath := "public/" + generateUniqueFilename(ext)

		// Open the uploaded file
		file, err := img.Open()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		defer file.Close()

		// Save image data to storage
		if err := c.storage.SaveImage(imagePath, file); err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		// Set the image path
		complaint.Images = append(complaint.Images, Image{Path: imagePath})
	}

	if err := c.UseCase.CreateComplaint(complaint); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, complaint)
}

func (c *ComplaintController) GetAllComplaint(ctx echo.Context) error {
	complaints, err := c.UseCase.GetAllComplaint()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, complaints)
}

func (c *ComplaintController) GetComplaintByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	complaint, err := c.UseCase.GetComplaintByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, complaint)
}

func (c *ComplaintController) UpdateComplaint(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	complaint := new(Complaint)
	if err := ctx.Bind(complaint); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.UseCase.UpdateComplaint(complaint, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, complaint)
}

func (c *ComplaintController) DeleteComplaint(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.UseCase.DeleteComplaint(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)

}

// Generate unique filename
func generateUniqueFilename(ext string) string {
	// Get current Unix timestamp in nanoseconds
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	// Replace characters that are not allowed in filenames
	// For example, replace ":" with "_" to ensure the filename is valid
	// You can add more replacements as needed
	// filename = strings.ReplaceAll(filename, ":", "_")

	// Combine timestamp and original filename to create a unique filename
	return timestamp + ext
}
