package user

type UseCase interface {
	RegisterUser(*User) error
	LoginUser(*User) (*User, error)
}

type userUseCase struct {
	repo Repository
}

func NewUserUseCase(repo Repository) UseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc *userUseCase) RegisterUser(user *User) error {
	return uc.repo.RegisterUser(user)
}

func (uc *userUseCase) LoginUser(user *User) (*User, error) {
	return uc.repo.LoginUser(user)
}
