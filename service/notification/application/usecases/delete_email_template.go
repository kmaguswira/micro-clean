package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/notification/application/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
)

type IDeleteEmailTemplate interface {
	Execute(ID string) (domain.EmailTemplate, error)
}

type deleteEmailTemplateUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewDeleteEmailTemplateUseCase(ReadWriteRepository repositories.IReadWriteRepository) IDeleteEmailTemplate {
	return &deleteEmailTemplateUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *deleteEmailTemplateUseCase) Execute(ID string) (domain.EmailTemplate, error) {
	result, err := t.readWriteRepository.DeleteEmailTemplate(ID)

	if err == nil {
		return *result, nil
	}

	return domain.EmailTemplate{}, errors.New("Bad Request")
}
