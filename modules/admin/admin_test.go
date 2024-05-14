package admin

import (
	"errors"
	"testing"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetEmailUser(email string) (*Admin, error) {
	args := m.Called(email)
	return args.Get(0).(*Admin), args.Error(1)
}

func (m *MockRepository) RegisterAdmin(admin *Admin) error {
	args := m.Called(admin)
	return args.Error(0)
}

func (m *MockRepository) LoginAdmin(admin *Admin) (*Admin, error) {
	args := m.Called(admin)
	return args.Get(0).(*Admin), args.Error(1)
}

func (m *MockRepository) UpdateStatusComplaint(id, status_id int, updatedAt time.Time) error {
	args := m.Called(id, status_id, updatedAt)
	return args.Error(0)
}

func (m *MockRepository) GetAllComplaint() ([]complaint.Complaint, error) {
	args := m.Called()
	return args.Get(0).([]complaint.Complaint), args.Error(1)
}

func (m *MockRepository) GetAllUser() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *MockRepository) UpdatePasswordUser(id int, pass string) error {
	args := m.Called(id, pass)
	return args.Error(0)
}

func (m *MockRepository) ActivateUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) IsActiveUser(id int) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) GetAllComplaintWithPaginate(page, perPage int) ([]complaint.Complaint, error) {
	args := m.Called(page, perPage)
	return args.Get(0).([]complaint.Complaint), args.Error(1)
}

func (m *MockRepository) getCountOfComplaints() int {
	args := m.Called()
	return args.Int(0)
}

func TestRegisterAdmin(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	admin := &Admin{
		Email:    "admin@example.com",
		Password: "password",
		Name:     "Admin",
	}

	mockRepo.On("GetEmailUser", admin.Email).Return((*Admin)(nil), errors.New("not found"))
	mockRepo.On("RegisterAdmin", admin).Return(nil)

	code, err := uc.RegisterAdmin(admin)
	assert.Equal(t, constants.SuccessCode, code)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestRegisterAdmin_EmailAlreadyExist(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	admin := &Admin{
		Email:    "admin@example.com",
		Password: "password",
		Name:     "Admin",
	}

	mockRepo.On("GetEmailUser", admin.Email).Return(admin, nil)

	code, err := uc.RegisterAdmin(admin)
	assert.Equal(t, constants.ErrCodeEmailAlreadyExist, code)
	assert.EqualError(t, err, constants.ErrEmailAlreadyExist)

	mockRepo.AssertExpectations(t)
}

func TestLoginAdmin(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	admin := &Admin{
		Email:    "admin@example.com",
		Password: "password",
	}

	mockRepo.On("LoginAdmin", admin).Return(admin, nil)

	resp, code, err := uc.LoginAdmin(admin)
	assert.Equal(t, constants.SuccessCode, code)
	assert.NoError(t, err)
	assert.Equal(t, admin, resp)

	mockRepo.AssertExpectations(t)
}

func TestUpdateStatusComplaint(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	id := 1
	status_id := 2

	mockRepo.On("UpdateStatusComplaint", id, status_id, mock.Anything).Return(nil)

	code, err := uc.UpdateStatusComplaint(id, status_id)
	assert.Equal(t, constants.SuccessCode, code)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllComplaint(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	complaints := []complaint.Complaint{{}, {}}

	mockRepo.On("GetAllComplaint").Return(complaints, nil)

	resp, code, err := uc.GetAllComplaint()
	assert.Equal(t, constants.SuccessCode, code)
	assert.NoError(t, err)
	assert.Equal(t, complaints, resp)

	mockRepo.AssertExpectations(t)
}

func TestGetAllUser(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	users := []user.User{{}, {}}

	mockRepo.On("GetAllUser").Return(users, nil)

	resp, code, err := uc.GetAllUser()
	assert.Equal(t, constants.SuccessCode, code)
	assert.NoError(t, err)
	assert.Equal(t, users, resp)

	mockRepo.AssertExpectations(t)
}

func TestUpdatePasswordUser(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	id := 1
	pass := "newpassword"

	mockRepo.On("UpdatePasswordUser", id, pass).Return(nil)

	code, err := uc.UpdatePasswordUser(id, pass)
	assert.Equal(t, constants.SuccessCode, code)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestActivateUser(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	id := 1

	mockRepo.On("IsActiveUser", id).Return(false, nil)
	mockRepo.On("ActivateUser", id).Return(nil)

	code, err := uc.ActivateUser(id)
	assert.Equal(t, constants.SuccessCode, code)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestActivateUser_AlreadyDeleted(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	id := 1

	mockRepo.On("IsActiveUser", id).Return(true, nil)

	code, err := uc.ActivateUser(id)
	assert.Equal(t, constants.ErrorCodeBadRequest, code)
	assert.EqualError(t, err, constants.ErrUserAlreadyDeleted)

	mockRepo.AssertExpectations(t)
}

func TestGetAllComplaintWithPaginate(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := NewAdminUseCase(mockRepo)

	page := 1
	perPage := 10
	complaints := []complaint.Complaint{{}, {}}
	totalCount := 20

	mockRepo.On("GetAllComplaintWithPaginate", page, perPage).Return(complaints, nil)
	mockRepo.On("getCountOfComplaints").Return(totalCount)

	resp, pagination, err := uc.GetAllComplaintWithPaginate(page, perPage)
	assert.NoError(t, err)
	assert.Equal(t, complaints, resp)
	assert.Equal(t, totalCount, pagination.TotalCount)
	assert.Equal(t, page, pagination.Page)
	assert.Equal(t, perPage, pagination.PerPage)

	mockRepo.AssertExpectations(t)
}
