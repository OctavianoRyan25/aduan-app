// internal/complaint/repo.go
package complaint

import "gorm.io/gorm"

type Repository interface {
	CreateComplaint(*Complaint) error
	GetAllComplaint(int) ([]Complaint, error)
	GetComplaintByID(int) (*Complaint, error)
	UpdateComplaint(*Complaint, int) error
	DeleteComplaint(int) error
	DeleteImageID(int) error
	GetImagesByComplaintID(int) ([]Image, error)
}

type complaintRepository struct {
	db *gorm.DB
}

func NewComplaintRepository(db *gorm.DB) Repository {
	return &complaintRepository{
		db: db,
	}
}

func (r *complaintRepository) CreateComplaint(complaint *Complaint) error {
	return r.db.Create(complaint).Error
}

func (r *complaintRepository) GetAllComplaint(user_id int) ([]Complaint, error) {
	var complaints []Complaint
	if err := r.db.Preload("Images").Preload("Status").Preload("User").Where("user_id = ?", user_id).Find(&complaints).Error; err != nil {
		return nil, err
	}
	return complaints, nil
}

func (r *complaintRepository) GetComplaintByID(id int) (*Complaint, error) {
	var complaint Complaint
	if err := r.db.Preload("Images").Preload("Status").Preload("User").Where("id = ?", id).First(&complaint).Error; err != nil {
		return nil, err
	}
	return &complaint, nil
}

func (r *complaintRepository) UpdateComplaint(complaint *Complaint, id int) error {
	return r.db.Model(&Complaint{}).Where("id = ?", id).Updates(complaint).Error
}

func (r *complaintRepository) DeleteComplaint(id int) error {
	return r.db.Delete(&Complaint{}, id).Error
}

func (r *complaintRepository) DeleteImageID(id int) error {
	return r.db.Delete(&Image{}, id).Error
}

func (r *complaintRepository) GetImagesByComplaintID(id int) ([]Image, error) {
	var images []Image
	if err := r.db.Where("complaint_id = ?", id).Find(&images).Error; err != nil {
		return nil, err // Handle error
	}
	return images, nil // Success
}
