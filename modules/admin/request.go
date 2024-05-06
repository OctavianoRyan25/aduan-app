package admin

type AdminRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
