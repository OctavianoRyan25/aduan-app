package user

type UserRegisterResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
