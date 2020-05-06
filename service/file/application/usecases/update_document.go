package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IUpdateDocument interface {
	Execute(input UpdateDocumentInput) (domain.Document, error)
}

type UpdateDocumentInput struct {
	ID          string
	Name        string
	Path        string
	Slug        string
	Type        string
	Description string
	Info        string
	Cdn         bool
}

type updateDocumentUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewUpdateDocumentUseCase(ReadWriteRepository repositories.IReadWriteRepository) IUpdateDocument {
	return &updateDocumentUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *updateDocumentUseCase) Execute(input UpdateDocumentInput) (domain.Document, error) {
	document := domain.Document{
		ID:          input.ID,
		Name:        input.Name,
		Path:        input.Path,
		Slug:        input.Slug,
		Type:        input.Type,
		Description: input.Description,
		Info:        input.Info,
		Cdn:         input.Cdn,
	}

	result, err := t.readWriteRepository.UpdateDocument(&document)

	if err == nil {
		return *result, nil
	}

	return domain.Document{}, errors.New("Bad Request")
}
