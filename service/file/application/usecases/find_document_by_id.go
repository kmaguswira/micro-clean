package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IFindImageByID interface {
	Execute(ID string) (domain.Image, error)
}

type findImageByIDUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindImageByIDUseCase(ReadRepository repositories.IReadRepository) IFindImageByID {
	return &findImageByIDUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findImageByIDUseCase) Execute(ID string) (domain.Image, error) {
	result, err := t.readRepository.FindImageByID(ID)

	if err == nil {
		return *result, nil
	}
	return domain.Image{}, errors.New("Bad Request")
}
