package admin

import (
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
)

func MapToComplaintResponse(admin Admin) AdminRegisterResponse {

	return AdminRegisterResponse{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
	}
}

func mapToComplaintResponse(complaint complaint.Complaint) ComplaintResponse {
	var imagesResponse []ImageResponse
	for _, img := range complaint.Images {
		imagesResponse = append(imagesResponse, ImageResponse{
			ID:   img.ID,
			Path: img.Path,
		})
	}

	//Mapping User to UserRegisterResponse
	mappingUser := user.MapToComplaintResponse(complaint.User)
	mappingStatus := mapStatusToResponse(complaint.Status)
	return ComplaintResponse{
		ID:         complaint.ID,
		Name:       complaint.Name,
		Phone:      complaint.Phone,
		Body:       complaint.Body,
		Category:   complaint.Category,
		Images:     imagesResponse,
		StatusID:   complaint.StatusID,
		Status:     mappingStatus,
		UserID:     complaint.UserID,
		User:       mappingUser,
		Location:   complaint.Location,
		Created_at: complaint.Created_at,
	}
}

func mapStatusToResponse(status complaint.Status) StatusResponse {
	return StatusResponse{
		ID:     status.ID,
		Status: status.Status,
	}
}

func MapToUserResponse(user user.User) UserResponse {

	return UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}
}
