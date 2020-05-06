package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IFindDocumentByID interface {
	Execute(ID string) (domain.Document, error)
}

type findDocumentByIDUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindDocumentByIDUseCase(ReadRepository repositories.IReadRepository) IFindDocumentByID {
	return &findDocumentByIDUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findDocumentByIDUseCase) Execute(ID string) (domain.Document, error) {
	result, err := t.readRepository.FindDocumentByID(ID)

	if err == nil {
		return *result, nil
	}
	return domain.Document{}, errors.New("Bad Request")
}
