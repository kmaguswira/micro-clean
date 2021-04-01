package handler

import (
	"context"

	. "github.com/ahmetb/go-linq"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/kmaguswira/micro-clean/service/file/application/global"
	"github.com/kmaguswira/micro-clean/service/file/application/usecases"
	"github.com/kmaguswira/micro-clean/service/file/domain"
	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
)

func (t *File) CreateImage(ctx context.Context, req *file.CreateImageRequest, res *file.CreateImageResponse) error {
	log.Log("Received File.CreateImage request")

	input := usecases.CreateImageInput{
		Name:        req.New.Name,
		Path:        req.New.Path,
		Slug:        req.New.Slug,
		Thumbnail:   req.New.Thumbnail,
		Type:        req.New.Type,
		Title:       req.New.Title,
		Alt:         req.New.Alt,
		Description: req.New.Description,
		Info:        req.New.Info,
		Cdn:         req.New.Cdn,
	}

	result, err := t.createImageUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	imageResponse := populateImageResponse(result)
	res.ResponseInfo = t.response.Created()
	res.Result = &imageResponse

	return nil
}

func (t *File) FindImageById(ctx context.Context, req *file.FindImageByIdRequest, res *file.FindImageByIdResponse) error {
	log.Log("Received File.FindImageById")

	result, err := t.findImageByIDUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	imageResponse := populateImageResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &imageResponse

	return nil
}

func (t *File) FindAllImage(ctx context.Context, req *file.FindAllImageRequest, res *file.FindAllImageResponse) error {
	log.Log("Received File.FindAllImage")

	input := global.FindAllInput{
		QueryKey: req.Query.QueryKey,
		Limit:    int(req.Query.Limit),
		Offset:   int(req.Query.Offset),
		Sort:     req.Query.Sort,
	}

	input.ParseValue(req.Query.QueryValue)

	result, err := t.findAllImageUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	var images []*file.Image
	From(result).Select(func(c interface{}) interface{} {
		d := c.(domain.Image)
		image := populateImageResponse(d)
		return &image
	}).ToSlice(&images)

	res.ResponseInfo = t.response.OK()
	res.Result = images
	return nil
}

func (t *File) UpdateImage(ctx context.Context, req *file.UpdateImageRequest, res *file.UpdateImageResponse) error {
	log.Log("Received File.UpdateImage")

	input := usecases.UpdateImageInput{
		ID:          req.Update.ID,
		Name:        req.Update.Name,
		Path:        req.Update.Path,
		Slug:        req.Update.Slug,
		Thumbnail:   req.Update.Thumbnail,
		Type:        req.Update.Type,
		Title:       req.Update.Title,
		Alt:         req.Update.Alt,
		Description: req.Update.Description,
		Info:        req.Update.Info,
		Cdn:         req.Update.Cdn,
	}

	result, err := t.updateImageUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	imageResponse := populateImageResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &imageResponse

	return nil
}

func (t *File) DeleteImage(ctx context.Context, req *file.DeleteImageRequest, res *file.DeleteImageResponse) error {
	log.Log("Received Notification.DeleteImage")

	result, err := t.deleteImageUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	imageResponse := populateImageResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &imageResponse

	return nil
}

func (t *File) CreateDocument(ctx context.Context, req *file.CreateDocumentRequest, res *file.CreateDocumentResponse) error {
	log.Log("Received File.CreateDocument request")

	input := usecases.CreateDocumentInput{
		Name:        req.New.Name,
		Path:        req.New.Path,
		Slug:        req.New.Slug,
		Type:        req.New.Type,
		Description: req.New.Description,
		Info:        req.New.Info,
		Cdn:         req.New.Cdn,
	}

	result, err := t.createDocumentUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	documentResponse := populateDocumentResponse(result)
	res.ResponseInfo = t.response.Created()
	res.Result = &documentResponse

	return nil
}

func (t *File) FindDocumentById(ctx context.Context, req *file.FindDocumentByIdRequest, res *file.FindDocumentByIdResponse) error {
	log.Log("Received File.FindDocumentById")

	result, err := t.findDocumentByIDUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	documentResponse := populateDocumentResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &documentResponse

	return nil
}

func (t *File) FindAllDocument(ctx context.Context, req *file.FindAllDocumentRequest, res *file.FindAllDocumentResponse) error {
	log.Log("Received File.FindAllDocument")

	input := global.FindAllInput{
		QueryKey: req.Query.QueryKey,
		Limit:    int(req.Query.Limit),
		Offset:   int(req.Query.Offset),
		Sort:     req.Query.Sort,
	}

	input.ParseValue(req.Query.QueryValue)

	result, err := t.findAllDocumentUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	var documents []*file.Document
	From(result).Select(func(c interface{}) interface{} {
		d := c.(domain.Document)
		document := populateDocumentResponse(d)
		return &document
	}).ToSlice(&documents)

	res.ResponseInfo = t.response.OK()
	res.Result = documents
	return nil
}

func (t *File) UpdateDocument(ctx context.Context, req *file.UpdateDocumentRequest, res *file.UpdateDocumentResponse) error {
	log.Log("Received File.UpdateDocument")

	input := usecases.UpdateDocumentInput{
		ID:          req.Update.ID,
		Name:        req.Update.Name,
		Path:        req.Update.Path,
		Slug:        req.Update.Slug,
		Type:        req.Update.Type,
		Description: req.Update.Description,
		Info:        req.Update.Info,
		Cdn:         req.Update.Cdn,
	}

	result, err := t.updateDocumentUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	documentResponse := populateDocumentResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &documentResponse

	return nil
}

func (t *File) DeleteDocument(ctx context.Context, req *file.DeleteDocumentRequest, res *file.DeleteDocumentResponse) error {
	log.Log("Received Notification.DeleteDocument")

	result, err := t.deleteDocumentUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	documentResponse := populateDocumentResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &documentResponse

	return nil
}
