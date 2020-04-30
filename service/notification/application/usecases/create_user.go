package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type ICreateUser interface {
	Execute(input CreateUserInput) (domain.User, error)
}

type CreateUserInput struct {
	Name     string
	Email    string
	Username string
	Status   string
	Password string
	RoleID   string
}

type createUserUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewCreateUserUseCase(ReadWriteRepository repositories.IReadWriteRepository) ICreateUser {
	return &createUserUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *createUserUseCase) Execute(input CreateUserInput) (domain.User, error) {
	newUser := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Status:   input.Status,
		Username: input.Username,
		RoleID:   input.RoleID,
	}

	newUser.SetPassword(input.Password)
	newUser.GenerateActivationToken()

	result, err := t.readWriteRepository.CreateUser(&newUser)

	if err == nil {
		return *result, nil
	}

	return domain.User{}, errors.New("Bad Request")
}
