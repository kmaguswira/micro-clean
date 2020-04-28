package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IForgotPassword interface {
	Execute(email string) (domain.User, error)
}

type forgotPasswordUseCase struct {
	readRepository      repositories.IReadRepository
	readWriteRepository repositories.IReadWriteRepository
}

func NewForgotPasswordUseCase(ReadRepository repositories.IReadRepository, ReadWriteRepository repositories.IReadWriteRepository) IForgotPassword {
	return &forgotPasswordUseCase{
		readRepository:      ReadRepository,
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *forgotPasswordUseCase) Execute(email string) (domain.User, error) {
	user, err := t.readRepository.FindUserByEmailOrUsername(email)

	if err != nil {
		return domain.User{}, errors.New("Bad Request")
	}

	user.GenerateForgotActivationToken()
	t.readWriteRepository.SetForgotPasswordToken(*user)

	return *user, nil
}
