package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/asim/go-micro/v3/client"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/rs/xid"
)

type uploadController struct {
	utils.Response
	fileService file.FileService
}

func NewUploadController(client client.Client) uploadController {
	return uploadController{
		fileService: file.NewFileService("kmaguswira.srv.file", client),
	}
}

func (t *uploadController) Images(c *gin.Context) {
	typeUpload := c.PostForm("type")
	titleUpload := c.PostForm("title")
	descriptionUpload := c.PostForm("description")
	var images []file.Image

	// Multipart form
	form, err := c.MultipartForm()

	if err != nil {
		t.BadRequest(c, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["files"]

	for _, f := range files {

		// build filename
		str := strings.Split(f.Filename, ".")
		fileType := str[len(str)-1]
		newFileName := xid.New().String() + "." + fileType

		// build path for image and thumbnails
		dir := filepath.Join("public", "images", typeUpload)
		dirLarge := filepath.Join(dir, "large")
		dirMedium := filepath.Join(dir, "medium")
		dirThumbnails := filepath.Join(dir, "thumbnails")
		dest := filepath.Join(dir, newFileName)

		// create folder
		os.MkdirAll(dir, os.ModePerm)
		os.MkdirAll(dirLarge, os.ModePerm)
		os.MkdirAll(dirMedium, os.ModePerm)
		os.MkdirAll(dirThumbnails, os.ModePerm)

		// save picture from form to the dest folder
		if err := c.SaveUploadedFile(f, dest); err != nil {
			t.BadRequest(c, fmt.Sprintf("get form err: %s", fmt.Sprintf("upload file err: %s", err.Error())))
			return
		}

		// open the image and then create the thumbnail
		src, _ := imaging.Open(dest)
		dstImage1920 := imaging.Resize(src, 1920, 0, imaging.Lanczos)
		largePath := filepath.Join(dirLarge, newFileName)
		dstImage800 := imaging.Resize(src, 800, 0, imaging.Lanczos)
		mediumPath := filepath.Join(dirMedium, newFileName)
		dstImage10 := imaging.Resize(src, 10, 10, imaging.Lanczos)
		thumbPath := filepath.Join(dirThumbnails, newFileName)

		if err = imaging.Save(dstImage1920, largePath); err != nil {
			t.BadRequest(c, fmt.Sprintf("get form err: %s", fmt.Sprintf("upload file err: %s", err.Error())))
			return
		}

		if err = imaging.Save(dstImage800, mediumPath); err != nil {
			t.BadRequest(c, fmt.Sprintf("get form err: %s", fmt.Sprintf("upload file err: %s", err.Error())))
			return
		}

		if err = imaging.Save(dstImage10, thumbPath); err != nil {
			t.BadRequest(c, fmt.Sprintf("get form err: %s", fmt.Sprintf("upload file err: %s", err.Error())))
			return
		}

		// save to file service
		response, err := t.fileService.CreateImage(c, &file.CreateImageRequest{
			New: &file.Image{
				Name:        newFileName,
				Path:        dir,
				Slug:        dest,
				Thumbnail:   thumbPath,
				Type:        typeUpload,
				Alt:         titleUpload + "__@__" + newFileName,
				Title:       titleUpload,
				Description: descriptionUpload,
				Info:        fmt.Sprintf("Filename: %s \nFileType: %s \nFileSize: %d", f.Filename, fileType, f.Size),
				Cdn:         false,
			},
		})

		if err != nil {
			t.BadRequest(c, err.Error())
			return
		}

		//TODO: publish event to create cdn

		images = append(images, *response.Result)
	}
	t.OKSingleData(c, images)

	return
}

func (t *uploadController) Documents(c *gin.Context) {
	typeUpload := c.PostForm("type")
	descriptionUpload := c.PostForm("description")
	var documents []file.Document

	// Multipart form
	form, err := c.MultipartForm()

	if err != nil {
		t.BadRequest(c, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["files"]

	for _, f := range files {

		// build filename
		str := strings.Split(f.Filename, ".")
		fileType := str[len(str)-1]
		newFileName := xid.New().String() + "." + fileType

		// build path for image and thumbnails
		dir := filepath.Join("public", "documents", typeUpload)
		dest := filepath.Join(dir, newFileName)

		// create folder
		os.MkdirAll(dir, os.ModePerm)

		// save document from form to the dest folder
		if err := c.SaveUploadedFile(f, dest); err != nil {
			t.BadRequest(c, fmt.Sprintf("get form err: %s", fmt.Sprintf("upload file err: %s", err.Error())))
			return
		}

		// save to file service
		response, err := t.fileService.CreateDocument(c, &file.CreateDocumentRequest{
			New: &file.Document{
				Name:        newFileName,
				Path:        dir,
				Slug:        dest,
				Type:        typeUpload,
				Description: descriptionUpload,
				Info:        fmt.Sprintf("Filename: %s \nFileType: %s \nFileSize: %d", f.Filename, fileType, f.Size),
				Cdn:         false,
			},
		})

		if err != nil {
			t.BadRequest(c, err.Error())
			return
		}

		//TODO: publish event to create cdn

		documents = append(documents, *response.Result)
	}
	t.OKSingleData(c, documents)

	return
}
