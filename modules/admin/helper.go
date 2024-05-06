package admin

func MapToComplaintResponse(admin Admin) AdminRegisterResponse {

	return AdminRegisterResponse{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
	}
}
