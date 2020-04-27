package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindAllRole interface {
	Execute(input global.FindAllInput) ([]domain.Role, error)
}

type findAllRoleUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindAllRoleUseCase(ReadRepository repositories.IReadRepository) IFindAllRole {
	return &findAllRoleUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findAllRoleUseCase) Execute(input global.FindAllInput) ([]domain.Role, error) {
	result, err := t.readRepository.FindAllRole(input)

	if err == nil {
		return *result, nil
	}
	return []domain.Role{}, errors.New("Bad Request")
}
