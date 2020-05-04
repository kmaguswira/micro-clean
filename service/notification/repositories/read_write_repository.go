package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	iface "github.com/kmaguswira/micro-clean/service/notification/application/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/config"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
	"github.com/kmaguswira/micro-clean/service/notification/repositories/entity"
)

type readWriteRepository struct {
	db *gorm.DB
}

func NewReadWriteRepository(DB *gorm.DB) iface.IReadWriteRepository {
	if DB != nil {
		return &readWriteRepository{
			db: DB,
		}
	}
	config := config.GetConfig()
	return &readWriteRepository{
		db: NewDB(config.Repositories.ReadWrite, Registered),
	}
}

func (t *readWriteRepository) CreateEmailTemplate(emailTemplate *domain.EmailTemplate) (*domain.EmailTemplate, error) {
	newEmailTemplate := entity.EmailTemplate{
		Title:    emailTemplate.Title,
		Subject:  emailTemplate.Subject,
		Body:     emailTemplate.Body,
		Language: emailTemplate.Language,
	}
	t.db.Create(&newEmailTemplate)

	emailTemplate.ID = newEmailTemplate.ID

	return emailTemplate, nil
}

func (t *readWriteRepository) DeleteEmailTemplate(ID string) (*domain.EmailTemplate, error) {
	var emailTemplate entity.EmailTemplate

	if err := t.db.Where("id = ?", ID).First(&emailTemplate).Error; err != nil {
		err := fmt.Errorf("EmailTemplate with ID %q not found", ID)
		return nil, err
	}

	emailTemplateDomain := populateEmailTemplateDomain(emailTemplate)
	t.db.Delete(&emailTemplate)

	return &emailTemplateDomain, nil
}

func (t *readWriteRepository) UpdateEmailTemplate(emailTemplateUpdated *domain.EmailTemplate) (*domain.EmailTemplate, error) {
	var emailTemplate entity.EmailTemplate

	if err := t.db.Where("id = ?", emailTemplateUpdated.ID).First(&emailTemplate).Error; err != nil {
		err := fmt.Errorf("EmailTemplate with ID %q not found", emailTemplateUpdated.ID)
		return nil, err
	}

	emailTemplate.Title = emailTemplateUpdated.Title
	emailTemplate.Subject = emailTemplateUpdated.Subject
	emailTemplate.Body = emailTemplateUpdated.Body
	emailTemplate.Language = emailTemplateUpdated.Language

	t.db.Save(&emailTemplate)
	return emailTemplateUpdated, nil
}
