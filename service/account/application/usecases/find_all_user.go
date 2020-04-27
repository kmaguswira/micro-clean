package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindAllUser interface {
	Execute(input global.FindAllInput) ([]domain.User, error)
}

type findAllUserUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindAllUserUseCase(ReadRepository repositories.IReadRepository) IFindAllUser {
	return &findAllUserUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findAllUserUseCase) Execute(input global.FindAllInput) ([]domain.User, error) {
	result, err := t.readRepository.FindAllUser(input)

	if err == nil {
		return *result, nil
	}
	return []domain.User{}, errors.New("Bad Request")
}
