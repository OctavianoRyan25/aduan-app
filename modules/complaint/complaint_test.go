package complaint

import (
	"errors"
	"testing"
	"time"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	CreateFn                 func(*Complaint) error
	GetAllFn                 func(int) ([]Complaint, error)
	GetByIDFn                func(int) (*Complaint, error)
	UpdateFn                 func(int, *Complaint) error
	DeleteFn                 func(int) error
	GetImagesByComplaintIDFn func(int) ([]Image, error)
	DeleteImageIDFn          func(int) error
}

func (m *mockRepo) CreateComplaint(cmp *Complaint) error {
	return m.CreateFn(cmp)
}

func (m *mockRepo) GetAllComplaint(userID int) ([]Complaint, error) {
	return m.GetAllFn(userID)
}

func (m *mockRepo) GetComplaintByID(id int) (*Complaint, error) {
	return m.GetByIDFn(id)
}

func (m *mockRepo) UpdateComplaint(cmp *Complaint, id int) error {
	return m.UpdateFn(id, cmp)
}

func (m *mockRepo) DeleteComplaint(id int) error {
	return m.DeleteFn(id)
}

func (m *mockRepo) GetImagesByComplaintID(id int) ([]Image, error) {
	return m.GetImagesByComplaintIDFn(id)
}

func (m *mockRepo) DeleteImageID(id int) error {
	return m.DeleteImageIDFn(id)
}

func TestCreateComplaint_Success(t *testing.T) {
	mock := &mockRepo{
		CreateFn: func(cmp *Complaint) error {
			return nil
		},
	}
	uc := NewComplaintUseCase(mock)

	cmp := &Complaint{
		Name:       "John Doe",
		Phone:      "123456789",
		Body:       "Test complaint",
		Category:   "Test category",
		Created_at: time.Now(),
		Updated_at: time.Now(),
		StatusID:   1,
	}

	code, err := uc.CreateComplaint(cmp)

	assert.NoError(t, err)
	assert.Equal(t, constants.SuccessCode, code)
}

func TestCreateComplaint_ValidationFailed(t *testing.T) {
	mock := &mockRepo{}
	uc := NewComplaintUseCase(mock)

	cmp := &Complaint{
		// Missing required fields intentionally to trigger validation error
	}

	code, err := uc.CreateComplaint(cmp)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrorCodeFieldRequired, code)
}

func TestGetAllComplaint_Success(t *testing.T) {
	mock := &mockRepo{
		GetAllFn: func(userID int) ([]Complaint, error) {
			// Mock data for successful retrieval
			complaints := []Complaint{
				{ID: 1, UserID: userID, Name: "John Doe", Phone: "123456789", Body: "Test complaint", Category: "Test category", Created_at: time.Now(), Updated_at: time.Now(), StatusID: 1},
				{ID: 2, UserID: userID, Name: "Jane Doe", Phone: "987654321", Body: "Another test complaint", Category: "Another test category", Created_at: time.Now(), Updated_at: time.Now(), StatusID: 1},
			}
			return complaints, nil
		},
	}
	uc := NewComplaintUseCase(mock)

	userID := 1 // Assuming a valid user ID
	complaints, code, err := uc.GetAllComplaint(userID)

	assert.NoError(t, err)
	assert.Equal(t, constants.SuccessCode, code)
	assert.Len(t, complaints, 2)
}

func TestGetAllComplaint_NotFound(t *testing.T) {
	mock := &mockRepo{
		GetAllFn: func(userID int) ([]Complaint, error) {
			// No complaints found
			return nil, errors.New(constants.ErrNotFound)
		},
	}
	uc := NewComplaintUseCase(mock)

	userID := 1 // Assuming a valid user ID
	complaints, code, err := uc.GetAllComplaint(userID)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeNotFound, code)
	assert.Nil(t, complaints)
}

func TestGetComplaintByID_Success(t *testing.T) {
	mock := &mockRepo{
		GetByIDFn: func(id int) (*Complaint, error) {
			// Mock data for successful retrieval
			complaintData := &Complaint{
				ID:         id,
				UserID:     1, // Assuming a valid user ID
				Name:       "John Doe",
				Phone:      "123456789",
				Body:       "Test complaint",
				Category:   "Test category",
				Created_at: time.Now(),
				Updated_at: time.Now(),
				StatusID:   1,
			}
			return complaintData, nil
		},
	}
	uc := NewComplaintUseCase(mock)

	id := 1     // Assuming a valid complaint ID
	userID := 1 // Assuming a valid user ID
	complaintData, code, err := uc.GetComplaintByID(id, userID)

	assert.NoError(t, err)
	assert.Equal(t, constants.SuccessCode, code)
	assert.NotNil(t, complaintData)
}

func TestGetComplaintByID_Unauthorized(t *testing.T) {
	mock := &mockRepo{
		GetByIDFn: func(id int) (*Complaint, error) {
			// Mock data for successful retrieval, but with different user ID
			complaintData := &Complaint{
				ID:         id,
				UserID:     2, // Different user ID
				Name:       "John Doe",
				Phone:      "123456789",
				Body:       "Test complaint",
				Category:   "Test category",
				Created_at: time.Now(),
				Updated_at: time.Now(),
				StatusID:   1,
			}
			return complaintData, nil
		},
	}
	uc := NewComplaintUseCase(mock)

	id := 1     // Assuming a valid complaint ID
	userID := 1 // Assuming a valid user ID
	complaintData, code, err := uc.GetComplaintByID(id, userID)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeUnauthorized, code)
	assert.Nil(t, complaintData)
}

// Test for UpdateComplaint function
func TestUpdateComplaint_Success(t *testing.T) {
	mock := &mockRepo{
		GetByIDFn: func(id int) (*Complaint, error) {
			// Mock data for existing complaint
			complaintData := &Complaint{
				ID:         id,
				UserID:     1, // Assuming a valid user ID
				Name:       "John Doe",
				Phone:      "123456789",
				Body:       "Test complaint",
				Category:   "Test category",
				Created_at: time.Now(),
				Updated_at: time.Now(),
				StatusID:   1,
			}
			return complaintData, nil
		},
		UpdateFn: func(id int, cmp *Complaint) error {
			// Mock update success
			return nil
		},
	}
	uc := NewComplaintUseCase(mock)

	id := 1     // Assuming a valid complaint ID
	userID := 1 // Assuming a valid user ID
	cmp := &Complaint{
		Name:     "Updated Name",
		Phone:    "987654321",
		Body:     "Updated complaint",
		Category: "Updated category",
	}
	code, err := uc.UpdateComplaint(id, userID, cmp)

	assert.NoError(t, err)
	assert.Equal(t, constants.SuccessCode, code)
}

// Test for DeleteComplaint function
func TestDeleteComplaint_Success(t *testing.T) {
	mock := &mockRepo{
		GetByIDFn: func(id int) (*Complaint, error) {
			// Mock data for existing complaint
			complaintData := &Complaint{
				ID:         id,
				UserID:     1, // Assuming a valid user ID
				Name:       "John Doe",
				Phone:      "123456789",
				Body:       "Test complaint",
				Category:   "Test category",
				Created_at: time.Now(),
				Updated_at: time.Now(),
				StatusID:   1,
			}
			return complaintData, nil
		},
		DeleteFn: func(id int) error {
			// Mock delete success
			return nil
		},
		GetImagesByComplaintIDFn: func(id int) ([]Image, error) {
			// Mock no images found for the complaint
			return nil, nil
		},
		DeleteImageIDFn: func(id int) error {
			// Mock delete image success
			return nil
		},
	}
	uc := NewComplaintUseCase(mock)

	id := 1     // Assuming a valid complaint ID
	userID := 1 // Assuming a valid user ID
	code, err := uc.DeleteComplaint(id, userID)

	assert.NoError(t, err)
	assert.Equal(t, constants.SuccessCode, code)
}

func TestGetComplaintByID_InvalidID(t *testing.T) {
	mock := &mockRepo{}
	uc := NewComplaintUseCase(mock)

	id := 0     // Invalid ID
	userID := 1 // Assuming a valid user ID
	complaintData, code, err := uc.GetComplaintByID(id, userID)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeInvalidID, code)
	assert.Nil(t, complaintData)
}

// Test for UpdateComplaint function with invalid ID
func TestUpdateComplaint_InvalidID(t *testing.T) {
	mock := &mockRepo{}
	uc := NewComplaintUseCase(mock)

	id := 0     // Invalid ID
	userID := 1 // Assuming a valid user ID
	cmp := &Complaint{
		Name:     "Updated Name",
		Phone:    "987654321",
		Body:     "Updated complaint",
		Category: "Updated category",
	}
	code, err := uc.UpdateComplaint(id, userID, cmp)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeInvalidID, code)
}

// Test for DeleteComplaint function with invalid ID
func TestDeleteComplaint_InvalidID(t *testing.T) {
	mock := &mockRepo{
		GetByIDFn: func(id int) (*Complaint, error) {
			// Mock data for complaint not found
			return nil, errors.New(constants.ErrNotFound)
		},
	}
	uc := NewComplaintUseCase(mock)

	id := 0     // Invalid ID
	userID := 1 // Assuming a valid user ID
	code, err := uc.DeleteComplaint(id, userID)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrCodeInvalidID, code)
}

func TestDeleteComplaint_WithImages_Success(t *testing.T) {
	// Setup mock repository with necessary functions
	mock := &mockRepo{
		GetByIDFn: func(id int) (*Complaint, error) {
			// Mock data for existing complaint with ID
			complaintData := &Complaint{
				ID:         id,
				UserID:     1, // Assuming a valid user ID
				Name:       "John Doe",
				Phone:      "123456789",
				Body:       "Test complaint",
				Category:   "Test category",
				Created_at: time.Now(),
				Updated_at: time.Now(),
				StatusID:   1,
			}
			return complaintData, nil
		},
		DeleteFn: func(id int) error {
			// Mock successful deletion
			return nil
		},
		GetImagesByComplaintIDFn: func(id int) ([]Image, error) {
			// Mock data for images attached to complaint
			images := []Image{
				{ID: 1, ComplaintID: id, Path: "http://example.com/image1.jpg"},
				{ID: 2, ComplaintID: id, Path: "http://example.com/image2.jpg"},
			}
			return images, nil
		},
		DeleteImageIDFn: func(id int) error {
			// Mock successful deletion of image
			return nil
		},
	}
	// Initialize complaint use case with mock repository
	uc := NewComplaintUseCase(mock)

	// Define valid complaint ID and user ID
	id := 1     // Assuming a valid complaint ID
	userID := 1 // Assuming a valid user ID

	// Perform deletion
	code, err := uc.DeleteComplaint(id, userID)

	// Assertion
	assert.NoError(t, err)                       // Expecting no error
	assert.Equal(t, constants.SuccessCode, code) // Expecting success code
}

func TestDeleteComplaint_ComplaintNotFound(t *testing.T) {
	// Setup mock repository with GetByIDFn returning error for complaint not found
	mock := &mockRepo{
		GetByIDFn: func(id int) (*Complaint, error) {
			// Mock error for complaint not found
			return nil, errors.New(constants.ErrNotFound)
		},
	}
	// Initialize complaint use case with mock repository
	uc := NewComplaintUseCase(mock)

	// Define valid complaint ID and user ID
	id := 1     // Assuming a valid complaint ID
	userID := 1 // Assuming a valid user ID

	// Perform deletion
	code, err := uc.DeleteComplaint(id, userID)

	// Assertion
	assert.Error(t, err)                             // Expecting an error
	assert.Equal(t, constants.ErrCodeNotFound, code) // Expecting complaint not found error code
}
