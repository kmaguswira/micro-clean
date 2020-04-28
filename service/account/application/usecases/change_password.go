package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IChangePassword interface {
	Execute(ID string, oldPassword string, newPassword string) (domain.User, error)
}

type changePasswordUseCase struct {
	readRepository      repositories.IReadRepository
	readWriteRepository repositories.IReadWriteRepository
}

func NewChangePasswordUseCase(ReadRepository repositories.IReadRepository, ReadWriteRepository repositories.IReadWriteRepository) IChangePassword {
	return &changePasswordUseCase{
		readRepository:      ReadRepository,
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *changePasswordUseCase) Execute(ID string, oldPassword string, newPassword string) (domain.User, error) {
	user, err := t.readRepository.FindUserByID(ID)

	if err != nil {
		return domain.User{}, errors.New("Bad Request")
	}

	if !user.CheckHashedPassword(oldPassword) {
		return domain.User{}, errors.New("Password doesnt match")
	}

	user.SetPassword(newPassword)
	t.readWriteRepository.ChangePassword(*user)

	return *user, nil
}
