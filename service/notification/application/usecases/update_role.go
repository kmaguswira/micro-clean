package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IUpdateRole interface {
	Execute(input UpdateRoleInput) (domain.Role, error)
}

type UpdateRoleInput struct {
	ID    string
	Title string
}

type updateRoleUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewUpdateRoleUseCase(ReadWriteRepository repositories.IReadWriteRepository) IUpdateRole {
	return &updateRoleUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *updateRoleUseCase) Execute(input UpdateRoleInput) (domain.Role, error) {
	role := domain.Role{
		ID:    input.ID,
		Title: input.Title,
	}

	result, err := t.readWriteRepository.UpdateRole(&role)

	if err == nil {
		return *result, nil
	}

	return domain.Role{}, errors.New("Bad Request")
}
