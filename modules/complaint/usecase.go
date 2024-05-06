package complaint

import (
	"errors"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
)

type UseCase interface {
	CreateComplaint(*Complaint) (int, error)
	GetAllComplaint(int) ([]Complaint, int, error)
	GetComplaintByID(id, user_id int) (*Complaint, int, error)
	UpdateComplaint(id, user_id int, complaint *Complaint) (int, error)
	DeleteComplaint(id, user_id int) (int, error)
}

type complaintUseCase struct {
	repo Repository
}

func NewComplaintUseCase(repo Repository) UseCase {
	return &complaintUseCase{
		repo: repo,
	}
}

func (uc *complaintUseCase) CreateComplaint(complaint *Complaint) (int, error) {
	//Validate complaint
	if complaint.Name == "" {
		return constants.ErrorCodeFieldRequired, errors.New(constants.ErrFieldRequired)
	}
	if complaint.Phone == "" {
		return constants.ErrorCodeFieldRequired, errors.New(constants.ErrFieldRequired)
	}
	if complaint.Body == "" {
		return constants.ErrorCodeFieldRequired, errors.New(constants.ErrFieldRequired)
	}
	if complaint.Category == "" {
		return constants.ErrorCodeFieldRequired, errors.New(constants.ErrFieldRequired)
	}
	complaint.Created_at = time.Now()
	complaint.Updated_at = time.Now()
	complaint.StatusID = 1
	err := uc.repo.CreateComplaint(complaint)
	if err != nil {
		return constants.ErrorCodeBadRequest, err
	}
	return 200, nil
}

func (uc *complaintUseCase) GetAllComplaint(user_id int) ([]Complaint, int, error) {
	complaints, err := uc.repo.GetAllComplaint(user_id)
	if err != nil {
		return nil, constants.ErrCodeNotFound, errors.New(constants.ErrNotFound)
	}
	return complaints, 200, err
}

func (uc *complaintUseCase) GetComplaintByID(id, user_id int) (*Complaint, int, error) {
	if id == 0 {
		return nil, constants.ErrCodeInvalidID, errors.New(constants.ErrInvalidID)
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return nil, constants.ErrCodeNotFound, err
	}

	if user_id != complaintdata.UserID {
		return nil, constants.ErrCodeUnauthorized, errors.New(constants.ErrUnauthorized)
	}

	return complaintdata, 200, err
}

func (uc *complaintUseCase) UpdateComplaint(id, user_id int, complaint *Complaint) (int, error) {
	if id == 0 {
		return constants.ErrCodeInvalidID, errors.New(constants.ErrInvalidID)
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return constants.ErrCodeNotFound, err
	}

	if user_id != complaintdata.UserID {
		return constants.ErrCodeUnauthorized, errors.New(constants.ErrUnauthorized)
	}

	complaint.ID = id
	complaint.UserID = user_id
	complaint.Updated_at = time.Now()
	err = uc.repo.UpdateComplaint(complaint, id)
	return 200, err
}

func (uc *complaintUseCase) DeleteComplaint(id, user_id int) (int, error) {
	if id == 0 {
		return constants.ErrCodeInvalidID, errors.New(constants.ErrInvalidID)
	}

	complaintdata, err := uc.repo.GetComplaintByID(id)
	if err != nil {
		return constants.ErrCodeNotFound, err
	}

	if user_id != complaintdata.UserID {
		return constants.ErrCodeUnauthorized, errors.New(constants.ErrUnauthorized)
	}

	images, err := uc.repo.GetImagesByComplaintID(id)
	if err != nil {
		return constants.ErrorCodeBadRequest, err
	}

	for _, image := range images {
		if err := uc.repo.DeleteImageID(image.ID); err != nil {
			return constants.ErrorCodeBadRequest, err
		}
	}

	return 200, uc.repo.DeleteComplaint(id)
}
