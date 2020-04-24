package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindRoleByID interface {
	Execute(ID string) (domain.Role, error)
}

type findRoleByIDUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindRoleByIDUseCase(ReadRepository repositories.IReadRepository) IFindRoleByID {
	return &findRoleByIDUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findRoleByIDUseCase) Execute(ID string) (domain.Role, error) {
	result, err := t.readRepository.FindRoleByID(ID)

	if err == nil {
		return *result, nil
	}
	return domain.Role{}, errors.New("Bad Request")
}
