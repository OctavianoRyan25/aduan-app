package admin

import (
	"errors"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
)

type UseCase interface {
	RegisterAdmin(*Admin) (int, error)
	LoginAdmin(*Admin) (*Admin, int, error)
	UpdateStatusComplaint(id, status_id int) (int, error)
	GetAllComplaint() ([]complaint.Complaint, int, error)
	GetAllUser() ([]user.User, int, error)
	UpdatePasswordUser(int, string) (int, error)
	ActivateUser(int) (int, error)
	GetAllComplaintWithPaginate(int, int) ([]complaint.Complaint, Pagination, error)
}

type adminUseCase struct {
	repo Repository
}

func NewAdminUseCase(repo Repository) UseCase {
	return &adminUseCase{
		repo: repo,
	}
}

func (uc *adminUseCase) RegisterAdmin(admin *Admin) (int, error) {
	if admin.Email == "" || admin.Password == "" || admin.Name == "" {
		return constants.ErrorCodeBadRequest, errors.New(constants.ErrFieldRequired)
	}
	email := admin.Email
	_, err := uc.repo.GetEmailUser(email)
	if err == nil {
		return constants.ErrCodeEmailAlreadyExist, errors.New(constants.ErrEmailAlreadyExist)
	}
	admin.Created_at = time.Now()
	admin.Updated_at = time.Now()
	return constants.SuccessCode, uc.repo.RegisterAdmin(admin)
}

func (uc *adminUseCase) LoginAdmin(admin *Admin) (*Admin, int, error) {
	if admin.Email == "" || admin.Password == "" {
		return nil, constants.ErrorCodeBadRequest, errors.New(constants.ErrFieldRequired)
	}
	resp, err := uc.repo.LoginAdmin(admin)
	return resp, constants.SuccessCode, err
}

func (uc *adminUseCase) UpdateStatusComplaint(id, status_id int) (int, error) {
	if id == 0 || status_id == 0 {
		return constants.ErrorCodeBadRequest, errors.New(constants.ErrFieldRequired)
	}
	updatedAt := time.Now()
	return constants.SuccessCode, uc.repo.UpdateStatusComplaint(id, status_id, updatedAt)
}

func (uc *adminUseCase) GetAllComplaint() ([]complaint.Complaint, int, error) {
	resp, err := uc.repo.GetAllComplaint()
	return resp, constants.SuccessCode, err
}

func (uc *adminUseCase) GetAllUser() ([]user.User, int, error) {
	resp, err := uc.repo.GetAllUser()
	return resp, constants.SuccessCode, err
}

func (uc *adminUseCase) UpdatePasswordUser(id int, pass string) (int, error) {
	if id == 0 || pass == "" {
		return constants.ErrorCodeBadRequest, errors.New(constants.ErrFieldRequired)
	}
	return constants.SuccessCode, uc.repo.UpdatePasswordUser(id, pass)
}

func (uc *adminUseCase) ActivateUser(id int) (int, error) {
	// Memeriksa apakah pengguna ada dan belum dihapus
	exists, err := uc.repo.IsActiveUser(id)
	if err != nil {
		return constants.ErrorCodeBadRequest, err
	}

	// Mengaktifkan pengguna
	if !exists {
		return constants.SuccessCode, uc.repo.ActivateUser(id)

	}
	// Jika pengguna tidak ditemukan atau sudah dihapus
	return constants.ErrorCodeBadRequest, errors.New(constants.ErrUserAlreadyDeleted)
}

func (uc *adminUseCase) GetAllComplaintWithPaginate(page, perPage int) ([]complaint.Complaint, Pagination, error) {
	resp, err := uc.repo.GetAllComplaintWithPaginate(page, perPage)
	if err != nil {
		return nil, Pagination{}, err
	}

	totalCount := uc.repo.getCountOfComplaints()

	pagination := NewPagination(page, perPage, totalCount)

	return resp, pagination, nil
}
