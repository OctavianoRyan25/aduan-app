package admin

import (
	"math"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
)

type AdminRegisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
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

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalCount int `json:"totalCount"`
	TotalPages int `json:"totalPages"`
}

type SuccessResponseWithPaginate struct {
	Status   string     `json:"status"`
	Message  string     `json:"message"`
	Data     any        `json:"data"`
	MetaData Pagination `json:"metaData"`
}

func NewPagination(page, perPage, totalCount int) Pagination {
	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))
	return Pagination{
		Page:       page,
		PerPage:    perPage,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}
}

type UserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
