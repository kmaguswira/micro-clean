package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IDeleteACL interface {
	Execute(ID string) (domain.ACL, error)
}

type deleteACLUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewDeleteACLUseCase(ReadWriteRepository repositories.IReadWriteRepository) IDeleteACL {
	return &deleteACLUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *deleteACLUseCase) Execute(ID string) (domain.ACL, error) {
	result, err := t.readWriteRepository.DeleteACL(ID)

	if err == nil {
		return *result, nil
	}

	return domain.ACL{}, errors.New("Bad Request")
}
