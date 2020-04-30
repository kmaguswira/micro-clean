package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type ICreateRole interface {
	Execute(input string) (domain.Role, error)
}

type createRoleUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewCreateRoleUseCase(ReadWriteRepository repositories.IReadWriteRepository) ICreateRole {
	return &createRoleUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *createRoleUseCase) Execute(input string) (domain.Role, error) {
	newRole := domain.Role{
		Title: input,
	}
	result, err := t.readWriteRepository.CreateRole(&newRole)

	if err == nil {
		return *result, nil
	}
	return domain.Role{}, errors.New("Bad Request")
}
