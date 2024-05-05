package complaint

import (
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
)

type Complaint struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Phone      string
	Body       string
	Category   string
	Images     []Image `gorm:"foreignKey:ComplaintID"`
	StatusID   int
	Status     Status `gorm:"foreignKey:StatusID"`
	UserID     int
	User       user.User `gorm:"foreignKey:UserID"`
	Location   string
	Created_at time.Time
	Updated_at time.Time
}

type Image struct {
	ID          int `gorm:"primaryKey"`
	ComplaintID int
	Path        string
}

type Status struct {
	ID     int `gorm:"primaryKey"`
	Status string
}
