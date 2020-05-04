package repositories

import (
	"github.com/kmaguswira/micro-clean/service/notification/domain"
)

type IReadWriteRepository interface {
	CreateEmailTemplate(emailTemplate *domain.EmailTemplate) (*domain.EmailTemplate, error)
	DeleteEmailTemplate(ID string) (*domain.EmailTemplate, error)
	UpdateEmailTemplate(emailTemplateUpdated *domain.EmailTemplate) (*domain.EmailTemplate, error)
}
