package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IDeleteDocument interface {
	Execute(ID string) (domain.Document, error)
}

type deleteDocumentUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewDeleteDocumentUseCase(ReadWriteRepository repositories.IReadWriteRepository) IDeleteDocument {
	return &deleteDocumentUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *deleteDocumentUseCase) Execute(ID string) (domain.Document, error) {
	result, err := t.readWriteRepository.DeleteDocument(ID)

	if err == nil {
		return *result, nil
	}

	return domain.Document{}, errors.New("Bad Request")
}
