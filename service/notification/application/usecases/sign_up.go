package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type ISignUp interface {
	Execute(input SignUpInput) (SignUpOutput, error)
}

type SignUpInput struct {
	Email    string
	Password string
	Name     string
	Username string
	Role     string
}

type SignUpOutput struct {
	User domain.User
}

type signUpUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewSignUpUseCase(ReadWriteRepository repositories.IReadWriteRepository) ISignUp {
	return &signUpUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (s *signUpUseCase) Execute(input SignUpInput) (SignUpOutput, error) {
	newUser := domain.User{
		Email:    input.Email,
		Name:     input.Name,
		Username: input.Username,
		RoleID:   input.Role,
		Status:   global.NOT_VERIFIED_USER_STATUS,
	}
	newUser.SetPassword(input.Password)
	newUser.GenerateActivationToken()
	result, err := s.readWriteRepository.CreateUser(&newUser)

	if err == nil {
		return SignUpOutput{
			User: *result,
		}, nil
	}
	return SignUpOutput{}, errors.New("Bad Request")
}
