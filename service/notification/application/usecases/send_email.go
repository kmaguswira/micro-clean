package usecases

import (
	"errors"
	"fmt"

	"github.com/kmaguswira/micro-clean/service/notification/application/external_service"
	"github.com/kmaguswira/micro-clean/service/notification/application/global"
	"github.com/kmaguswira/micro-clean/service/notification/application/repositories"
)

type ISendEmail interface {
	Execute(templateTitle string, toName string, toEmail string, data []interface{}) (bool, error)
}

type sendEmailUseCase struct {
	readRepository   repositories.IReadRepository
	sendEmailService external_service.ISendEmail
}

func NewSendEmailUseCase(ReadRepository repositories.IReadRepository, SendEmailService external_service.ISendEmail) ISendEmail {
	return &sendEmailUseCase{
		readRepository:   ReadRepository,
		sendEmailService: SendEmailService,
	}
}

func (t *sendEmailUseCase) Execute(templateTitle string, toName string, toEmail string, data []interface{}) (bool, error) {
	result, err := t.readRepository.FindEmailTemplateByTitle(templateTitle)

	if err != nil {
		return false, errors.New("Bad Request")
	}

	input := global.SendEmailInput{
		To: global.EmailAddress{
			Name:  toName,
			Email: toEmail,
		},
		From: global.EmailAddress{
			Name:  result.FromName,
			Email: result.FromEmail,
		},
		Subject:          result.Subject,
		PlainTextContent: result.Subject,
		HTML:             fmt.Sprintf(result.HTML, data...),
	}

	response, err := t.sendEmailService.SendEmail(input)

	if err != nil {
		return false, errors.New("Failed to send email")
	}

	return response, nil
}
