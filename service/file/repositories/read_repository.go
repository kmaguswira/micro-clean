package repositories

import (
	"fmt"

	. "github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	"github.com/kmaguswira/micro-clean/service/file/application/global"
	iface "github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/config"
	"github.com/kmaguswira/micro-clean/service/file/domain"
	"github.com/kmaguswira/micro-clean/service/file/repositories/entity"
)

type readRepository struct {
	db *gorm.DB
}

func NewReadRepository(DB *gorm.DB) iface.IReadRepository {
	if DB != nil {
		return &readRepository{
			db: DB,
		}
	}
	config := config.GetConfig()
	return &readRepository{
		db: NewDB(config.Repositories.Read, Registered),
	}
}

func (t *readRepository) FindImageByID(input string) (*domain.Image, error) {
	var image entity.Image

	if err := t.db.Where("id = ?", input).First(&image).Error; err != nil {
		err := fmt.Errorf("Image Template with id %q not found", input)
		return nil, err
	}

	imageDomain := populateImageDomain(image)

	return &imageDomain, nil
}

func (t *readRepository) FindAllImages(input global.FindAllInput) (*[]domain.Image, error) {
	var images []entity.Image
	var result []domain.Image

	if err := t.db.Order(input.Sort).Limit(input.Limit).Offset(input.Offset).Where(input.QueryKey, input.QueryValue...).Find(&images).Error; err != nil {
		err := fmt.Errorf("Query Error", input)
		return nil, err
	}

	From(images).Select(func(c interface{}) interface{} {
		e := c.(entity.Image)
		return populateImageDomain(e)
	}).ToSlice(&result)

	return &result, nil
}

func (t *readRepository) FindDocumentByID(input string) (*domain.Document, error) {
	var document entity.Document

	if err := t.db.Where("id = ?", input).First(&document).Error; err != nil {
		err := fmt.Errorf("Document Template with id %q not found", input)
		return nil, err
	}

	DocumentDomain := populateDocumentDomain(document)

	return &DocumentDomain, nil
}

func (t *readRepository) FindAllDocuments(input global.FindAllInput) (*[]domain.Document, error) {
	var documents []entity.Document
	var result []domain.Document

	if err := t.db.Order(input.Sort).Limit(input.Limit).Offset(input.Offset).Where(input.QueryKey, input.QueryValue...).Find(&documents).Error; err != nil {
		err := fmt.Errorf("Query Error", input)
		return nil, err
	}

	From(documents).Select(func(c interface{}) interface{} {
		e := c.(entity.Document)
		return populateDocumentDomain(e)
	}).ToSlice(&result)

	return &result, nil
}
