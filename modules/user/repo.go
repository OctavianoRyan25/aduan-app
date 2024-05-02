package user

import "gorm.io/gorm"

type Repository interface {
	RegisterUser(*User) error
	LoginUser(*User) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) RegisterUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) LoginUser(user *User) (*User, error) {
	var u User
	if err := r.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
