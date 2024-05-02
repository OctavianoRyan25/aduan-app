// internal/complaint/usecase.go
package complaint

import (
	"errors"
	"time"
)

type UseCase interface {
	CreateComplaint(*Complaint) error
	GetAllComplaint() ([]Complaint, error)
	GetComplaintByID(int) (*Complaint, error)
	UpdateComplaint(*Complaint, int) error
	DeleteComplaint(int) error
}

type complaintUseCase struct {
	repo Repository
}

func NewComplaintUseCase(repo Repository) UseCase {
	return &complaintUseCase{
		repo: repo,
	}
}

func (uc *complaintUseCase) CreateComplaint(complaint *Complaint) error {
	complaint.Created_at = time.Now()
	complaint.Updated_at = time.Now()
	return uc.repo.CreateComplaint(complaint)
}

func (uc *complaintUseCase) GetAllComplaint() ([]Complaint, error) {
	return uc.repo.GetAllComplaint()
}

func (uc *complaintUseCase) GetComplaintByID(id int) (*Complaint, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return uc.repo.GetComplaintByID(id)
}

func (uc *complaintUseCase) UpdateComplaint(complaint *Complaint, id int) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	complaint.ID = id
	complaint.Updated_at = time.Now()
	return uc.repo.UpdateComplaint(complaint, id)
}

func (uc *complaintUseCase) DeleteComplaint(id int) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return uc.repo.DeleteComplaint(id)
}
