package repositories

import (
	"fmt"

	. "github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	"github.com/kmaguswira/micro-clean/service/notification/application/global"
	iface "github.com/kmaguswira/micro-clean/service/notification/application/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/config"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
	"github.com/kmaguswira/micro-clean/service/notification/repositories/entity"
)

type readRepository struct {
	db *gorm.DB
}

func NewReadRepository(DB *gorm.DB) iface.IReadRepository {
	if DB != nil {
		return &readRepository{
			db: DB,
		}
	}
	config := config.GetConfig()
	return &readRepository{
		db: NewDB(config.Repositories.Read, Registered),
	}
}

func (t *readRepository) FindEmailTemplateByID(input string) (*domain.EmailTemplate, error) {
	var emailTemplate entity.EmailTemplate

	if err := t.db.Where("id = ?", input).First(&emailTemplate).Error; err != nil {
		err := fmt.Errorf("Email Template with id %q not found", input)
		return nil, err
	}

	emailTemplateDomain := populateEmailTemplateDomain(emailTemplate)

	return &emailTemplateDomain, nil
}

func (t *readRepository) FindAllEmailTemplates(input global.FindAllInput) (*[]domain.EmailTemplate, error) {
	var emailTemplates []entity.EmailTemplate
	var result []domain.EmailTemplate

	if err := t.db.Order(input.Sort).Limit(input.Limit).Offset(input.Offset).Where(input.QueryKey, input.QueryValue...).Find(&emailTemplates).Error; err != nil {
		err := fmt.Errorf("Query Error", input)
		return nil, err
	}

	From(emailTemplates).Select(func(c interface{}) interface{} {
		e := c.(entity.EmailTemplate)
		return populateEmailTemplateDomain(e)
	}).ToSlice(&result)

	return &result, nil
}

func (t *readRepository) FindEmailTemplateByTitle(input string) (*domain.EmailTemplate, error) {
	var emailTemplate entity.EmailTemplate

	if err := t.db.Where("title = ?", input).First(&emailTemplate).Error; err != nil {
		err := fmt.Errorf("Email Template with title %q not found", input)
		return nil, err
	}

	emailTemplateDomain := populateEmailTemplateDomain(emailTemplate)

	return &emailTemplateDomain, nil
}
