package repositories

import (
	"github.com/kmaguswira/micro-clean/service/notification/application/global"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
)

type IReadRepository interface {
	FindEmailTemplateByID(input string) (*domain.EmailTemplate, error)
	FindAllEmailTemplates(input global.FindAllInput) (*[]domain.EmailTemplate, error)
	FindEmailTemplateByTitle(input string) (*domain.EmailTemplate, error)
}
