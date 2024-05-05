package user

import (
	"time"
)

type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Phone      string
	Address    string
	Created_at time.Time
	Updated_at time.Time
}

type Admin struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}
