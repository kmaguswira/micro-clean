package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/notification/application/global"
	"github.com/kmaguswira/micro-clean/service/notification/application/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
)

type IFindAllEmailTemplate interface {
	Execute(input global.FindAllInput) ([]domain.EmailTemplate, error)
}

type findAllEmailTemplateUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindAllEmailTemplateUseCase(ReadRepository repositories.IReadRepository) IFindAllEmailTemplate {
	return &findAllEmailTemplateUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findAllEmailTemplateUseCase) Execute(input global.FindAllInput) ([]domain.EmailTemplate, error) {
	result, err := t.readRepository.FindAllEmailTemplates(input)

	if err == nil {
		return *result, nil
	}
	return []domain.EmailTemplate{}, errors.New("Bad Request")
}
