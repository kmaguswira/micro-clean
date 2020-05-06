package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IDeleteImage interface {
	Execute(ID string) (domain.Image, error)
}

type deleteImageUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewDeleteImageUseCase(ReadWriteRepository repositories.IReadWriteRepository) IDeleteImage {
	return &deleteImageUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *deleteImageUseCase) Execute(ID string) (domain.Image, error) {
	result, err := t.readWriteRepository.DeleteImage(ID)

	if err == nil {
		return *result, nil
	}

	return domain.Image{}, errors.New("Bad Request")
}
