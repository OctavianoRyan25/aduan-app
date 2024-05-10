package admin

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}
type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Phone      string
	Address    string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
}
