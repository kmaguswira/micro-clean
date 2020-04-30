package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IUpdateUser interface {
	Execute(input UpdateUserInput) (domain.User, error)
}

type UpdateUserInput struct {
	ID       string
	Name     string
	Username string
	Email    string
	RoleID   string
	Status   string
}

type updateUserUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewUpdateUserUseCase(ReadWriteRepository repositories.IReadWriteRepository) IUpdateUser {
	return &updateUserUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *updateUserUseCase) Execute(input UpdateUserInput) (domain.User, error) {
	user := domain.User{
		ID:       input.ID,
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		RoleID:   input.RoleID,
		Status:   input.Status,
	}

	result, err := t.readWriteRepository.UpdateUser(&user)

	if err == nil {
		return *result, nil
	}

	return domain.User{}, errors.New("Bad Request")
}
