package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type ICreateDocument interface {
	Execute(input CreateDocumentInput) (domain.Document, error)
}

type CreateDocumentInput struct {
	Name        string
	Path        string
	Slug        string
	Type        string
	Description string
	Info        string
	Cdn         bool
}

type createDocumentUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewCreateDocumentUseCase(ReadWriteRepository repositories.IReadWriteRepository) ICreateDocument {
	return &createDocumentUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *createDocumentUseCase) Execute(input CreateDocumentInput) (domain.Document, error) {
	newDocument := domain.Document{
		Name:        input.Name,
		Path:        input.Path,
		Slug:        input.Slug,
		Type:        input.Type,
		Description: input.Description,
		Info:        input.Info,
		Cdn:         input.Cdn,
	}
	result, err := t.readWriteRepository.CreateDocument(&newDocument)

	if err == nil {
		return *result, nil
	}
	return domain.Document{}, errors.New("Bad Request")
}
