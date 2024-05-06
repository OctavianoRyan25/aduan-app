package user

import "gorm.io/gorm"

type Repository interface {
	RegisterUser(*User) error
	LoginUser(*User) (*User, error)
	GetEmailUser(email string) (*User, error)
	GetIDUser(id int) (*User, error)
	InactiveUser(*User) error
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
	if err := r.db.Where("email = ?", user.Email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetEmailUser(email string) (*User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetIDUser(id int) (*User, error) {
	var u User
	if err := r.db.Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) InactiveUser(user *User) error {
	return r.db.Where("email = ?", user.Email).Delete(user).Error
}
