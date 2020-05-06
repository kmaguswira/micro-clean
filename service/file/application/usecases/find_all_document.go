package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/global"
	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IFindAllImage interface {
	Execute(input global.FindAllInput) ([]domain.Image, error)
}

type findAllImageUseCase struct {
	readRepository repositories.IReadRepository
}

func NewFindAllImageUseCase(ReadRepository repositories.IReadRepository) IFindAllImage {
	return &findAllImageUseCase{
		readRepository: ReadRepository,
	}
}

func (t *findAllImageUseCase) Execute(input global.FindAllInput) ([]domain.Image, error) {
	result, err := t.readRepository.FindAllImages(input)

	if err == nil {
		return *result, nil
	}
	return []domain.Image{}, errors.New("Bad Request")
}
