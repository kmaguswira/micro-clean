package repositories

import (
	"github.com/kmaguswira/micro-clean/service/file/domain"
	"github.com/kmaguswira/micro-clean/service/file/repositories/entity"
)

func populateImageDomain(image entity.Image) domain.Image {
	imageDomain := domain.Image{
		ID:          image.ID,
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

	return imageDomain
}

func populateDocumentDomain(document entity.Document) domain.Document {
	documentDomain := domain.Document{
		ID:          document.ID,
		Name:        document.Name,
		Path:        document.Path,
		Slug:        document.Slug,
		Type:        document.Type,
		Description: document.Description,
		Info:        document.Info,
		Cdn:         document.Cdn,
	}

	return documentDomain
}
