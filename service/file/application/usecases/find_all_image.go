package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/global"
	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IFindAllDocument interface {
	Execute(input global.FindAllInput) ([]domain.Document, error)
}

type findAllDocumentUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindAllDocumentUseCase(ReadRepository repositories.IReadRepository) IFindAllDocument {
	return &findAllDocumentUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findAllDocumentUseCase) Execute(input global.FindAllInput) ([]domain.Document, error) {
	result, err := t.readRepository.FindAllDocuments(input)

	if err == nil {
		return *result, nil
	}
	return []domain.Document{}, errors.New("Bad Request")
}
