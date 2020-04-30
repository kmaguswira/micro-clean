package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IDeleteUser interface {
	Execute(ID string) (domain.User, error)
}

type deleteUserUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewDeleteUserUseCase(ReadWriteRepository repositories.IReadWriteRepository) IDeleteUser {
	return &deleteUserUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *deleteUserUseCase) Execute(ID string) (domain.User, error) {
	result, err := t.readWriteRepository.DeleteUser(ID)

	if err == nil {
		return *result, nil
	}

	return domain.User{}, errors.New("Bad Request")
}
