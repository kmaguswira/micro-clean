package controllers

import (
	"github.com/asim/go-micro/v3/client"
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
	"github.com/kmaguswira/micro-clean/utils"
)

type documentController struct {
	utils.Response
	fileService file.FileService
}

func NewDocumentController(client client.Client) documentController {
	return documentController{
		fileService: file.NewFileService("kmaguswira.srv.file", client),
	}
}

func (t *documentController) Create(c *gin.Context) {
	var createDocumentRequest requests.CreateDocument

	if err := c.BindJSON(&createDocumentRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.fileService.CreateDocument(c, &file.CreateDocumentRequest{
		New: &file.Document{
			Name:        createDocumentRequest.Name,
			Path:        createDocumentRequest.Path,
			Slug:        createDocumentRequest.Slug,
			Type:        createDocumentRequest.Type,
			Description: createDocumentRequest.Description,
			Info:        createDocumentRequest.Info,
			Cdn:         createDocumentRequest.Cdn,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t *documentController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := t.fileService.FindDocumentById(c, &file.FindDocumentByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *documentController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := t.fileService.DeleteDocument(c, &file.DeleteDocumentRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *documentController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateDocumentRequest requests.UpdateDocument

	if err := c.BindJSON(&updateDocumentRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.fileService.UpdateDocument(c, &file.UpdateDocumentRequest{
		Update: &file.Document{
			ID:          id,
			Name:        updateDocumentRequest.Name,
			Path:        updateDocumentRequest.Path,
			Slug:        updateDocumentRequest.Slug,
			Type:        updateDocumentRequest.Type,
			Description: updateDocumentRequest.Description,
			Info:        updateDocumentRequest.Info,
			Cdn:         updateDocumentRequest.Cdn,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *documentController) FindAll(c *gin.Context) {
	q := utils.QueryBuilder(c)

	response, err := t.fileService.FindAllDocument(c, &file.FindAllDocumentRequest{
		Query: &file.BaseQuery{
			QueryKey:   q.QueryKey,
			QueryValue: q.QueryValue,
			Limit:      q.Limit,
			Offset:     q.Offset,
			Sort:       q.Sort,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}
