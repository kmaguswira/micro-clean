package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/notification/application/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
)

type IFindEmailTemplateByID interface {
	Execute(ID string) (domain.EmailTemplate, error)
}

type findEmailTemplateByIDUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindEmailTemplateByIDUseCase(ReadRepository repositories.IReadRepository) IFindEmailTemplateByID {
	return &findEmailTemplateByIDUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findEmailTemplateByIDUseCase) Execute(ID string) (domain.EmailTemplate, error) {
	result, err := t.readRepository.FindEmailTemplateByID(ID)

	if err == nil {
		return *result, nil
	}
	return domain.EmailTemplate{}, errors.New("Bad Request")
}
