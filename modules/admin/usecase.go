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
	UpdatePasswordUser(*user.User) (int, error)
	ActivateUser(int) (int, error)
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
	admin.Created_at = time.Now()
	admin.Updated_at = time.Now()
	return constants.SuccessCode, uc.repo.RegisterAdmin(admin)
}

func (uc *adminUseCase) LoginAdmin(admin *Admin) (*Admin, int, error) {
	resp, err := uc.repo.LoginAdmin(admin)
	return resp, constants.SuccessCode, err
}

func (uc *adminUseCase) UpdateStatusComplaint(id, status_id int) (int, error) {
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

func (uc *adminUseCase) UpdatePasswordUser(user *user.User) (int, error) {
	return constants.SuccessCode, uc.repo.UpdatePasswordUser(user)
}

func (uc *adminUseCase) ActivateUser(id int) (int, error) {
	// Memeriksa apakah pengguna ada dan belum dihapus
	exists, err := uc.repo.IsActiveUser(id)
	if err != nil {
		return constants.ErrorCodeBadRequest, err
	}

	// Jika pengguna tidak ditemukan atau sudah dihapus
	if !exists {
		return constants.SuccessCode, uc.repo.ActivateUser(id)

	}
	// Mengaktifkan pengguna
	return constants.ErrorCodeBadRequest, errors.New("user not found or already deleted")
}
