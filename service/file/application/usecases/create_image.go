package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type ICreateImage interface {
	Execute(input CreateImageInput) (domain.Image, error)
}

type CreateImageInput struct {
	Name        string
	Path        string
	Slug        string
	Thumbnail   string
	Type        string
	Title       string
	Alt         string
	Description string
	Info        string
	Cdn         bool
}

type createImageUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewCreateImageUseCase(ReadWriteRepository repositories.IReadWriteRepository) ICreateImage {
	return &createImageUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *createImageUseCase) Execute(input CreateImageInput) (domain.Image, error) {
	newImage := domain.Image{
		Name:        input.Name,
		Path:        input.Path,
		Slug:        input.Slug,
		Thumbnail:   input.Thumbnail,
		Type:        input.Type,
		Title:       input.Title,
		Alt:         input.Alt,
		Description: input.Description,
		Info:        input.Info,
		Cdn:         input.Cdn,
	}
	result, err := t.readWriteRepository.CreateImage(&newImage)

	if err == nil {
		return *result, nil
	}
	return domain.Image{}, errors.New("Bad Request")
}
