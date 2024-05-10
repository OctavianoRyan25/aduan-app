package admin

import (
	"errors"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterAdmin(*Admin) error
	LoginAdmin(*Admin) (*Admin, error)
	UpdateStatusComplaint(id int, status_id int, updated_at time.Time) error
	GetAllComplaint() ([]complaint.Complaint, error)
	GetAllUser() ([]user.User, error)
	UpdatePasswordUser(int, string) error
	ActivateUser(int) error
	IsActiveUser(int) (bool, error)
	GetEmailUser(string) (*Admin, error)
	GetAllComplaintWithPaginate(int, int) ([]complaint.Complaint, error)
	getCountOfComplaints() int
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) Repository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) RegisterAdmin(admin *Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) LoginAdmin(admin *Admin) (*Admin, error) {
	var u Admin
	if err := r.db.Where("email = ?", admin.Email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *adminRepository) UpdateStatusComplaint(id int, status_id int, updated_at time.Time) error {
	err := r.db.Model(&complaint.Complaint{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status_id":  status_id,
		"updated_at": updated_at,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepository) GetAllComplaint() ([]complaint.Complaint, error) {
	var complaints []complaint.Complaint
	if err := r.db.Preload("Images").Preload("Status").Preload("User").Find(&complaints).Error; err != nil {
		return nil, err
	}
	return complaints, nil
}

func (r *adminRepository) GetAllUser() ([]user.User, error) {
	var users []user.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *adminRepository) UpdatePasswordUser(id int, pass string) error {
	var user user.User
	err := r.db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"password":   pass,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) ActivateUser(id int) error {
	//return r.db.Model(&user.User{}).Where("id = ?", id).Update("deleted_at", nil).Error
	query := "UPDATE users SET deleted_at = NULL WHERE id = ?"
	result := r.db.Exec(query, id).Error
	return result
}

func (r *adminRepository) IsActiveUser(id int) (bool, error) {
	var user user.User

	err := r.db.Unscoped().Where("id = ?", id).First(&user).Error

	if err != nil {
		return false, err
	}

	if user.Deleted_at.Time.IsZero() {
		return true, errors.New(constants.ErrUserAlreadyActive)
	}

	return false, nil
}

func (r *adminRepository) GetEmailUser(email string) (*Admin, error) {
	var a Admin
	if err := r.db.Where("email = ?", email).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *adminRepository) GetAllComplaintWithPaginate(page, perPage int) ([]complaint.Complaint, error) {
	var complaints []complaint.Complaint
	offset := (page - 1) * perPage

	if err := r.db.Preload("Images").
		Preload("Status").
		Preload("User").
		Offset(offset).
		Limit(perPage).
		Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}

func (r *adminRepository) getCountOfComplaints() int {
	var count int64
	r.db.Model(&complaint.Complaint{}).Count(&count)
	return int(count)
}
