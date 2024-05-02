package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Phone      string         `json:"phone"`
	Address    string         `json:"address"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
