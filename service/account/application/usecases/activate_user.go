package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IActivateUser interface {
	Execute(token string) (domain.User, error)
}

type activateUserUseCase struct {
	readRepository      repositories.IReadRepository
	readWriteRepository repositories.IReadWriteRepository
}

func NewActivateUserUseCase(ReadRepository repositories.IReadRepository, ReadWriteRepository repositories.IReadWriteRepository) IActivateUser {
	return &activateUserUseCase{
		readRepository:      ReadRepository,
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *activateUserUseCase) Execute(token string) (domain.User, error) {
	user, err := t.readRepository.FindUserByActivationToken(token)

	if err != nil {
		return domain.User{}, errors.New("Bad Request")
	}

	if user.Status == "pending" {
		user, _ := t.readWriteRepository.ActivateUser(user.ID)
		return *user, nil
	}

	return *user, nil
}
