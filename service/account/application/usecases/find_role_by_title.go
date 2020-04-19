package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindRoleByTitle interface {
	Execute(input string) (domain.Role, error)
}

type findRoleByTitleUseCase struct {
	readRepository repositories.IReadRepository
}

func NewfindRoleByTitleUseCase(ReadRepository repositories.IReadRepository) IFindRoleByTitle {
	return &findRoleByTitleUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findRoleByTitleUseCase) Execute(input string) (domain.Role, error) {

	result, err := t.readRepository.FindRoleByTitle(input)

	if err == nil {
		return *result, nil
	}
	return domain.Role{}, errors.New("Bad Request")
}
