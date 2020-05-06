package repositories

import (
	"github.com/kmaguswira/micro-clean/service/file/domain"
)

type IReadWriteRepository interface {
	CreateImage(image *domain.Image) (*domain.Image, error)
	DeleteImage(ID string) (*domain.Image, error)
	UpdateImage(imageUpdated *domain.Image) (*domain.Image, error)
	CreateDocument(document *domain.Document) (*domain.Document, error)
	DeleteDocument(ID string) (*domain.Document, error)
	UpdateDocument(documentUpdated *domain.Document) (*domain.Document, error)
}
