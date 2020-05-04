package repositories

import (
	"github.com/kmaguswira/micro-clean/service/notification/domain"
	"github.com/kmaguswira/micro-clean/service/notification/repositories/entity"
)

func populateEmailTemplateDomain(emailTemplate entity.EmailTemplate) domain.EmailTemplate {
	emailTemplateDomain := domain.EmailTemplate{
		ID:       emailTemplate.ID,
		Title:    emailTemplate.Title,
		Subject:  emailTemplate.Subject,
		Body:     emailTemplate.Body,
		Language: emailTemplate.Language,
	}

	return emailTemplateDomain
}
