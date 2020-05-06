package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro/client"
)

type imageController struct {
	utils.Response
	fileService file.FileService
}

func NewImageController(client client.Client) imageController {
	return imageController{
		fileService: file.NewFileService("kmaguswira.srv.file", client),
	}
}

func (t *imageController) Create(c *gin.Context) {
	var createImageRequest requests.CreateImage

	if err := c.BindJSON(&createImageRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.fileService.CreateImage(c, &file.CreateImageRequest{
		New: &file.Image{
			Name:        createImageRequest.Name,
			Path:        createImageRequest.Path,
			Slug:        createImageRequest.Slug,
			Thumbnail:   createImageRequest.Thumbnail,
			Type:        createImageRequest.Type,
			Title:       createImageRequest.Title,
			Alt:         createImageRequest.Alt,
			Description: createImageRequest.Description,
			Info:        createImageRequest.Info,
			Cdn:         createImageRequest.Cdn,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t *imageController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := t.fileService.FindImageById(c, &file.FindImageByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *imageController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := t.fileService.DeleteImage(c, &file.DeleteImageRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *imageController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateImageRequest requests.UpdateImage

	if err := c.BindJSON(&updateImageRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.fileService.UpdateImage(c, &file.UpdateImageRequest{
		Update: &file.Image{
			ID:          id,
			Name:        updateImageRequest.Name,
			Path:        updateImageRequest.Path,
			Slug:        updateImageRequest.Slug,
			Thumbnail:   updateImageRequest.Thumbnail,
			Type:        updateImageRequest.Type,
			Title:       updateImageRequest.Title,
			Alt:         updateImageRequest.Alt,
			Description: updateImageRequest.Description,
			Info:        updateImageRequest.Info,
			Cdn:         updateImageRequest.Cdn,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *imageController) FindAll(c *gin.Context) {
	q := utils.QueryBuilder(c)

	response, err := t.fileService.FindAllImage(c, &file.FindAllImageRequest{
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
