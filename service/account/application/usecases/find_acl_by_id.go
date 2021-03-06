package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindACLByID interface {
	Execute(ID string) (domain.ACL, error)
}

type findACLByIDUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindACLByIDUseCase(ReadRepository repositories.IReadRepository) IFindACLByID {
	return &findACLByIDUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findACLByIDUseCase) Execute(ID string) (domain.ACL, error) {
	result, err := t.readRepository.FindACLByID(ID)

	if err == nil {
		return *result, nil
	}

	return domain.ACL{}, errors.New("Bad Request")
}
