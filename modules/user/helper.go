package user

func MapToComplaintResponse(user User) UserRegisterResponse {

	return UserRegisterResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}
}
