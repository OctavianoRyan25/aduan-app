package complaint

import (
	"errors"
	"time"
)

type UseCase interface {
	CreateComplaint(*Complaint) error
	GetAllComplaint(int) ([]Complaint, error)
	GetComplaintByID(id, user_id int) (*Complaint, error)
	UpdateComplaint(id, user_id int, complaint *Complaint) error
	DeleteComplaint(id, user_id int) error
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
	complaint.StatusID = 1
	return uc.repo.CreateComplaint(complaint)
}

func (uc *complaintUseCase) GetAllComplaint(user_id int) ([]Complaint, error) {
	return uc.repo.GetAllComplaint(user_id)
}

func (uc *complaintUseCase) GetComplaintByID(id, user_id int) (*Complaint, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return nil, err
	}

	if user_id != complaintdata.UserID {
		return nil, errors.New("unauthorized")
	}

	return uc.repo.GetComplaintByID(id)
}

func (uc *complaintUseCase) UpdateComplaint(id, user_id int, complaint *Complaint) error {
	if id == 0 {
		return errors.New("invalid id")
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return err
	}

	if complaintdata == nil {
		return errors.New("complaint not found")
	}

	if user_id != complaintdata.UserID {
		return errors.New("unauthorized")
	}

	complaint.ID = id
	complaint.UserID = user_id
	complaint.Updated_at = time.Now()
	return uc.repo.UpdateComplaint(complaint, id)
}

func (uc *complaintUseCase) DeleteComplaint(id, user_id int) error {
	if id == 0 {
		return errors.New("invalid id")
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return err
	}

	if complaintdata == nil {
		return errors.New("complaint not found")
	}

	if user_id != complaintdata.UserID {
		return errors.New("unauthorized")
	}

	images, err := uc.repo.GetImagesByComplaintID(id)
	if err != nil {
		return err
	}

	for _, image := range images {
		if err := uc.repo.DeleteImageID(image.ID); err != nil {
			return err
		}
	}

	return uc.repo.DeleteComplaint(id)
}
