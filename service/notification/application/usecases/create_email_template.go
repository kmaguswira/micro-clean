package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/notification/application/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
)

type ICreateEmailTemplate interface {
	Execute(input CreateEmailTemplateInput) (domain.EmailTemplate, error)
}

type CreateEmailTemplateInput struct {
	Title    string
	Subject  string
	Body     string
	Language string
}

type createEmailTemplateUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewCreateEmailTemplateUseCase(ReadWriteRepository repositories.IReadWriteRepository) ICreateEmailTemplate {
	return &createEmailTemplateUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *createEmailTemplateUseCase) Execute(input CreateEmailTemplateInput) (domain.EmailTemplate, error) {
	newEmailTemplate := domain.EmailTemplate{
		Title:    input.Title,
		Subject:  input.Subject,
		Body:     input.Body,
		Language: input.Language,
	}
	result, err := t.readWriteRepository.CreateEmailTemplate(&newEmailTemplate)

	if err == nil {
		return *result, nil
	}
	return domain.EmailTemplate{}, errors.New("Bad Request")
}
