package complaint

import (
	"strconv"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
)

func generateUniqueFilename(ext string) string {
	// Mengubah nama file menjadi unix
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return timestamp + ext
}

func mapToCreateComplaintResponse(complaint Complaint) CreateComplaintResponse {
	var imagesResponse []ImageResponse
	for _, img := range complaint.Images {
		imagesResponse = append(imagesResponse, ImageResponse{
			ID:   img.ID,
			Path: img.Path,
		})
	}

	return CreateComplaintResponse{
		ID:         complaint.ID,
		Name:       complaint.Name,
		Phone:      complaint.Phone,
		Body:       complaint.Body,
		Category:   complaint.Category,
		Images:     imagesResponse,
		StatusID:   complaint.StatusID,
		UserID:     complaint.UserID,
		Location:   complaint.Location,
		Created_at: complaint.Created_at,
	}
}

func mapToComplaintResponse(complaint Complaint) ComplaintResponse {
	var imagesResponse []ImageResponse
	for _, img := range complaint.Images {
		imagesResponse = append(imagesResponse, ImageResponse{
			ID:   img.ID,
			Path: img.Path,
		})
	}

	//Mapping User to UserRegisterResponse
	mappingUser := user.MapToComplaintResponse(complaint.User)

	return ComplaintResponse{
		ID:         complaint.ID,
		Name:       complaint.Name,
		Phone:      complaint.Phone,
		Body:       complaint.Body,
		Category:   complaint.Category,
		Images:     imagesResponse,
		StatusID:   complaint.StatusID,
		Status:     complaint.Status,
		UserID:     complaint.UserID,
		User:       mappingUser,
		Location:   complaint.Location,
		Created_at: complaint.Created_at,
	}
}
