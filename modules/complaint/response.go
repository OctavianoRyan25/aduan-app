package complaint

import (
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
)

type CreateComplaintResponse struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Phone      string          `json:"phone"`
	Body       string          `json:"body"`
	Category   string          `json:"category"`
	Images     []ImageResponse `json:"images" `
	StatusID   int
	UserID     int
	Location   string    `json:"location"`
	Created_at time.Time `json:"created_at"`
}

type ComplaintResponse struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Phone      string          `json:"phone"`
	Body       string          `json:"body"`
	Category   string          `json:"category"`
	Images     []ImageResponse `json:"images" `
	StatusID   int
	Status     Status `gorm:"foreignKey:StatusID"`
	UserID     int
	User       user.UserRegisterResponse `gorm:"foreignKey:UserID"`
	Location   string                    `json:"location"`
	Created_at time.Time                 `json:"created_at"`
}

type ImageResponse struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Path string `json:"path"`
}
