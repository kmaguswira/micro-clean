package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindUserByID interface {
	Execute(id string) (domain.User, error)
}

type findUserByIDUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindUserByIDUseCase(ReadRepository repositories.IReadRepository) IFindUserByID {
	return &findUserByIDUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findUserByIDUseCase) Execute(id string) (domain.User, error) {
	result, err := t.readRepository.FindUserByID(id)

	if err == nil {
		return *result, nil
	}
	return domain.User{}, errors.New("Bad Request")
}
