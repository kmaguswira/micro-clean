package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindAllACL interface {
	Execute(input global.FindAllInput) ([]domain.ACL, error)
}

type findAllACLUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindAllACLUseCase(ReadRepository repositories.IReadRepository) IFindAllACL {
	return &findAllACLUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findAllACLUseCase) Execute(input global.FindAllInput) ([]domain.ACL, error) {
	result, err := t.readRepository.FindAllACL(input)

	if err == nil {
		return *result, nil
	}
	return []domain.ACL{}, errors.New("Bad Request")
}
