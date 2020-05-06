package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	iface "github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/config"
	"github.com/kmaguswira/micro-clean/service/file/domain"
	"github.com/kmaguswira/micro-clean/service/file/repositories/entity"
)

type readWriteRepository struct {
	db *gorm.DB
}

func NewReadWriteRepository(DB *gorm.DB) iface.IReadWriteRepository {
	if DB != nil {
		return &readWriteRepository{
			db: DB,
		}
	}
	config := config.GetConfig()
	return &readWriteRepository{
		db: NewDB(config.Repositories.ReadWrite, Registered),
	}
}

func (t *readWriteRepository) CreateImage(image *domain.Image) (*domain.Image, error) {
	newImage := entity.Image{
		Name:        image.Name,
		Path:        image.Path,
		Slug:        image.Slug,
		Thumbnail:   image.Thumbnail,
		Type:        image.Type,
		Title:       image.Title,
		Alt:         image.Alt,
		Description: image.Description,
		Info:        image.Info,
		Cdn:         image.Cdn,
	}
	t.db.Create(&newImage)

	image.ID = newImage.ID

	return image, nil
}

func (t *readWriteRepository) DeleteImage(ID string) (*domain.Image, error) {
	var image entity.Image

	if err := t.db.Where("id = ?", ID).First(&image).Error; err != nil {
		err := fmt.Errorf("Image with ID %q not found", ID)
		return nil, err
	}

	imageDomain := populateImageDomain(image)
	t.db.Delete(&image)

	return &imageDomain, nil
}

func (t *readWriteRepository) UpdateImage(imageUpdated *domain.Image) (*domain.Image, error) {
	var image entity.Image

	if err := t.db.Where("id = ?", imageUpdated.ID).First(&image).Error; err != nil {
		err := fmt.Errorf("Image with ID %q not found", imageUpdated.ID)
		return nil, err
	}

	image.Name = imageUpdated.Name
	image.Path = imageUpdated.Path
	image.Slug = imageUpdated.Slug
	image.Thumbnail = imageUpdated.Thumbnail
	image.Type = imageUpdated.Type
	image.Title = imageUpdated.Title
	image.Alt = imageUpdated.Alt
	image.Description = imageUpdated.Description
	image.Info = imageUpdated.Info
	image.Cdn = imageUpdated.Cdn

	t.db.Save(&image)
	return imageUpdated, nil
}

func (t *readWriteRepository) CreateDocument(document *domain.Document) (*domain.Document, error) {
	newDocument := entity.Document{
		Name:        document.Name,
		Path:        document.Path,
		Slug:        document.Slug,
		Type:        document.Type,
		Description: document.Description,
		Info:        document.Info,
		Cdn:         document.Cdn,
	}
	t.db.Create(&newDocument)

	document.ID = newDocument.ID

	return document, nil
}

func (t *readWriteRepository) DeleteDocument(ID string) (*domain.Document, error) {
	var document entity.Document

	if err := t.db.Where("id = ?", ID).First(&document).Error; err != nil {
		err := fmt.Errorf("Document with ID %q not found", ID)
		return nil, err
	}

	documentDomain := populateDocumentDomain(document)
	t.db.Delete(&document)

	return &documentDomain, nil
}

func (t *readWriteRepository) UpdateDocument(documentUpdated *domain.Document) (*domain.Document, error) {
	var document entity.Document

	if err := t.db.Where("id = ?", documentUpdated.ID).First(&document).Error; err != nil {
		err := fmt.Errorf("Document with ID %q not found", documentUpdated.ID)
		return nil, err
	}

	document.Name = documentUpdated.Name
	document.Path = documentUpdated.Path
	document.Slug = documentUpdated.Slug
	document.Type = documentUpdated.Type
	document.Description = documentUpdated.Description
	document.Info = documentUpdated.Info
	document.Cdn = documentUpdated.Cdn

	t.db.Save(&document)
	return documentUpdated, nil
}
