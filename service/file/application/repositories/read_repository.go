package repositories

import (
	"github.com/kmaguswira/micro-clean/service/file/application/global"
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IReadRepository interface {
	FindImageByID(input string) (*domain.Image, error)
	FindAllImages(input global.FindAllInput) (*[]domain.Image, error)
	FindDocumentByID(input string) (*domain.Document, error)
	FindAllDocuments(input global.FindAllInput) (*[]domain.Document, error)
}
