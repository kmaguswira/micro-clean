package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type ISignIn interface {
	Execute(input SignInInput) (SignInOutput, error)
}

type SignInInput struct {
	User     string
	Password string
}

type SignInOutput struct {
	User domain.User
}

type signInUseCase struct {
	readRepository repositories.IReadRepository
}

func NewSignInUseCase(ReadRepository repositories.IReadRepository) ISignIn {
	return &signInUseCase{
		readRepository: ReadRepository,
	}
}

func (t *signInUseCase) Execute(input SignInInput) (SignInOutput, error) {
	result, err := t.readRepository.FindUserByEmailOrUsername(input.User)

	if err != nil || !result.CheckHashedPassword(input.Password) {
		return SignInOutput{}, errors.New("Not Authorized")
	}

	return SignInOutput{
		User: *result,
	}, nil
}
