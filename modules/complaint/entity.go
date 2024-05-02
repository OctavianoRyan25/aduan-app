package complaint

import (
	"time"

	"gorm.io/gorm"
)

type Complaint struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Phone      string         `json:"phone"`
	Body       string         `json:"body"`
	Category   string         `json:"category"`
	Images     []Image        `json:"images" gorm:"foreignKey:ComplaintID"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Image struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ComplaintID int    `json:"complaint_id"`
	Path        string `json:"path"`
}

type ComplaintRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Body     string `json:"body"`
	Category string `json:"category"`
}
