package usecases

import (
	"errors"

	"github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IUpdateImage interface {
	Execute(input UpdateImageInput) (domain.Image, error)
}

type UpdateImageInput struct {
	ID          string
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

type updateImageUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewUpdateImageUseCase(ReadWriteRepository repositories.IReadWriteRepository) IUpdateImage {
	return &updateImageUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *updateImageUseCase) Execute(input UpdateImageInput) (domain.Image, error) {
	image := domain.Image{
		ID:          input.ID,
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

	result, err := t.readWriteRepository.UpdateImage(&image)

	if err == nil {
		return *result, nil
	}

	return domain.Image{}, errors.New("Bad Request")
}
