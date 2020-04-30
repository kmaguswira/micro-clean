package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IFindACLByHandler interface {
	Execute(handler string) (domain.ACL, error)
}

type findACLByHandlerUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindACLByHandlerUseCase(ReadRepository repositories.IReadRepository) IFindACLByHandler {
	return &findACLByHandlerUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findACLByHandlerUseCase) Execute(handler string) (domain.ACL, error) {
	result, err := t.readRepository.FindACLByHandler(handler)

	if err == nil {
		return *result, nil
	}

	return domain.ACL{}, errors.New("Bad Request")
}
