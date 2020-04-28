package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IResetPassword interface {
	Execute(token string, password string) (domain.User, error)
}

type resetPasswordUseCase struct {
	readRepository      repositories.IReadRepository
	readWriteRepository repositories.IReadWriteRepository
}

func NewResetPasswordUseCase(ReadRepository repositories.IReadRepository, ReadWriteRepository repositories.IReadWriteRepository) IResetPassword {
	return &resetPasswordUseCase{
		readRepository:      ReadRepository,
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *resetPasswordUseCase) Execute(token string, password string) (domain.User, error) {
	user, err := t.readRepository.FindUserByResetPasswordToken(token)

	if err != nil {
		return domain.User{}, errors.New("Bad Request")
	}

	user.SetForgotPasswordToken("")
	user.SetPassword(password)
	t.readWriteRepository.ResetPassword(*user)

	return *user, nil
}
