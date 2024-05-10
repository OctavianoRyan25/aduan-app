package user

import (
	"errors"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
)

type UseCase interface {
	RegisterUser(*User) (int, error)
	LoginUser(*User) (*User, int, error)
	InactiveUser(int, int) (int, error)
}

type userUseCase struct {
	repo Repository
}

func NewUserUseCase(repo Repository) UseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc *userUseCase) RegisterUser(user *User) (int, error) {
	email := user.Email
	_, err := uc.repo.GetEmailUser(email)
	if err == nil {
		return constants.ErrCodeEmailAlreadyExist, errors.New(constants.ErrEmailAlreadyExist)
	}
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	return 200, uc.repo.RegisterUser(user)
}

func (uc *userUseCase) LoginUser(user *User) (*User, int, error) {
	resp, err := uc.repo.LoginUser(user)
	return resp, 200, err
}

func (uc *userUseCase) InactiveUser(id, user_id int) (int, error) {
	user, err := uc.repo.GetIDUser(id)
	if err != nil {
		return constants.ErrCodeNotFound, err
	}

	if user == nil {
		return constants.ErrCodeNotFound, errors.New("user not found")
	}

	if user_id != user.ID {
		return constants.ErrCodeUnauthorized, errors.New(constants.ErrUnauthorized)
	}

	return 200, uc.repo.InactiveUser(user)
}
