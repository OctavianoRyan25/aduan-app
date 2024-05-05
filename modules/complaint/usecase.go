package complaint

import (
	"errors"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
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
	complaints, err := uc.repo.GetAllComplaint(user_id)
	if err != nil {
		return nil, errors.New(constants.ErrNotFound)
	}
	return complaints, err
}

func (uc *complaintUseCase) GetComplaintByID(id, user_id int) (*Complaint, error) {
	if id == 0 {
		return nil, errors.New(constants.ErrInvalidID)
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return nil, err
	}

	if user_id != complaintdata.UserID {
		return nil, errors.New(constants.ErrUnauthorized)
	}

	return uc.repo.GetComplaintByID(id)
}

func (uc *complaintUseCase) UpdateComplaint(id, user_id int, complaint *Complaint) error {
	if id == 0 {
		return errors.New(constants.ErrInvalidID)
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return err
	}

	if complaintdata == nil {
		return errors.New(constants.ErrNotFound)
	}

	if user_id != complaintdata.UserID {
		return errors.New(constants.ErrUnauthorized)
	}

	complaint.ID = id
	complaint.UserID = user_id
	complaint.Updated_at = time.Now()
	return uc.repo.UpdateComplaint(complaint, id)
}

func (uc *complaintUseCase) DeleteComplaint(id, user_id int) error {
	if id == 0 {
		return errors.New(constants.ErrInvalidID)
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return err
	}

	if complaintdata == nil {
		return errors.New(constants.ErrNotFound)
	}

	if user_id != complaintdata.UserID {
		return errors.New(constants.ErrUnauthorized)
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
