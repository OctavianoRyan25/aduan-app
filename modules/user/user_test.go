package user

import (
	"errors"
	"testing"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	GetEmailUserFn func(email string) (*User, error)
	RegisterUserFn func(user *User) error
	LoginUserFn    func(user *User) (*User, error)
	GetIDUserFn    func(id int) (*User, error)
	InactiveUserFn func(user *User) error
}

func (m *mockRepo) GetEmailUser(email string) (*User, error) {
	return m.GetEmailUserFn(email)
}

func (m *mockRepo) RegisterUser(user *User) error {
	return m.RegisterUserFn(user)
}

func (m *mockRepo) LoginUser(user *User) (*User, error) {
	return m.LoginUserFn(user)
}

func (m *mockRepo) GetIDUser(id int) (*User, error) {
	return m.GetIDUserFn(id)
}

func (m *mockRepo) InactiveUser(user *User) error {
	return m.InactiveUserFn(user)
}

func TestRegisterUser_Success(t *testing.T) {
	mock := &mockRepo{
		GetEmailUserFn: func(email string) (*User, error) {
			// No existing user with the same email
			return nil, nil
		},
		RegisterUserFn: func(user *User) error {
			// Mock successful registration
			return nil
		},
	}
	uc := NewUserUseCase(mock)

	userData := &User{
		Email:      "test@example.com",
		Password:   "password",
		Phone:      "08123456789",
		Address:    "Jl. Test No. 1",
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	code, err := uc.RegisterUser(userData)

	assert.NoError(t, err)
	assert.Equal(t, 200, code)
}

func TestRegisterUser_EmailAlreadyExist(t *testing.T) {
	mock := &mockRepo{
		GetEmailUserFn: func(email string) (*User, error) {
			// Existing user with the same email
			return &User{}, nil
		},
	}
	uc := NewUserUseCase(mock)

	userData := &User{
		Email: "test@example.com",
	}

	code, err := uc.RegisterUser(userData)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeEmailAlreadyExist, code)
}

func TestLoginUser_Success(t *testing.T) {
	mock := &mockRepo{
		LoginUserFn: func(user *User) (*User, error) {
			// Mock successful login
			return &User{
				ID:       1,
				Email:    "test@example.com",
				Password: "password",
			}, nil
		},
	}
	uc := NewUserUseCase(mock)

	userData := &User{
		Email:    "test@example.com",
		Password: "password",
	}

	resp, code, err := uc.LoginUser(userData)

	assert.NoError(t, err)
	assert.Equal(t, 200, code)
	assert.NotNil(t, resp)
}

func TestInactiveUser_Success(t *testing.T) {
	mock := &mockRepo{
		GetIDUserFn: func(id int) (*User, error) {
			// Mock user found
			return &User{
				ID: 1,
			}, nil
		},
		InactiveUserFn: func(user *User) error {
			// Mock successful inactive user
			return nil
		},
	}
	uc := NewUserUseCase(mock)

	code, err := uc.InactiveUser(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, 200, code)
}

func TestInactiveUser_UserNotFound(t *testing.T) {
	mock := &mockRepo{
		GetIDUserFn: func(id int) (*User, error) {
			// Mock user not found
			return nil, errors.New("user not found")
		},
	}
	uc := NewUserUseCase(mock)

	code, err := uc.InactiveUser(1, 1)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeNotFound, code)
}

func TestInactiveUser_Unauthorized(t *testing.T) {
	mock := &mockRepo{
		GetIDUserFn: func(id int) (*User, error) {
			// Mock user found, but not the same as the logged in user
			return &User{
				ID: 2,
			}, nil
		},
	}
	uc := NewUserUseCase(mock)

	code, err := uc.InactiveUser(1, 1)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeUnauthorized, code)
}
