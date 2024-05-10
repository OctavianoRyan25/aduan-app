package complaint

type ComplaintRequest struct {
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Body      string  `json:"body"`
	Category  string  `json:"category"`
	StatusID  int     `json:"status_id"`
	UserID    int     `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ImageRequest struct {
	Path string `json:"path"`
}

type LocationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
