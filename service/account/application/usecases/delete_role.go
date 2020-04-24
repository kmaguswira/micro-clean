package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IDeleteRole interface {
	Execute(ID string) (domain.Role, error)
}

type deleteRoleUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewDeleteRoleUseCase(ReadWriteRepository repositories.IReadWriteRepository) IDeleteRole {
	return &deleteRoleUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *deleteRoleUseCase) Execute(ID string) (domain.Role, error) {
	result, err := t.readWriteRepository.DeleteRole(ID)

	if err == nil {
		return *result, nil
	}

	return domain.Role{}, errors.New("Bad Request")
}
