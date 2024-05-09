package admin

import (
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
)

type AdminRegisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ComplaintResponse struct {
	ID         int                       `json:"id"`
	Name       string                    `json:"name"`
	Phone      string                    `json:"phone"`
	Body       string                    `json:"body"`
	Category   string                    `json:"category"`
	Images     []ImageResponse           `json:"images" `
	StatusID   int                       `json:"status_id"`
	Status     StatusResponse            `json:"status" gorm:"foreignKey:StatusID"`
	UserID     int                       `json:"user_id"`
	User       user.UserRegisterResponse `json:"user" gorm:"foreignKey:UserID"`
	Location   string                    `json:"location"`
	Created_at time.Time                 `json:"created_at"`
}

type ImageResponse struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Path string `json:"path"`
}

type StatusResponse struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}
