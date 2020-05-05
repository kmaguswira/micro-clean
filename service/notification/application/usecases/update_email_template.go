package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/notification/application/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
)

type IUpdateEmailTemplate interface {
	Execute(input UpdateEmailTemplateInput) (domain.EmailTemplate, error)
}

type UpdateEmailTemplateInput struct {
	ID        string
	Title     string
	Subject   string
	HTML      string
	PlainText string
	Language  string
	FromName  string
	FromEmail string
}

type updateEmailTemplateUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewUpdateEmailTemplateUseCase(ReadWriteRepository repositories.IReadWriteRepository) IUpdateEmailTemplate {
	return &updateEmailTemplateUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *updateEmailTemplateUseCase) Execute(input UpdateEmailTemplateInput) (domain.EmailTemplate, error) {
	emailTemplate := domain.EmailTemplate{
		ID:        input.ID,
		Title:     input.Title,
		Subject:   input.Subject,
		HTML:      input.HTML,
		PlainText: input.PlainText,
		Language:  input.Language,
		FromName:  input.FromName,
		FromEmail: input.FromEmail,
	}

	result, err := t.readWriteRepository.UpdateEmailTemplate(&emailTemplate)

	if err == nil {
		return *result, nil
	}

	return domain.EmailTemplate{}, errors.New("Bad Request")
}
