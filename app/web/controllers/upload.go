package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro/client"
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

	// Multipart form
	form, err := c.MultipartForm()

	if err != nil {
		t.BadRequest(c, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["files"]

	for _, file := range files {

		// build filename
		str := strings.Split(file.Filename, ".")
		fileType := str[len(str)-1]
		newFileName := xid.New().String() + "." + fileType

		// build path for image and thumbnails
		dir := filepath.Join("public", "images", typeUpload)
		dirThumbnails := filepath.Join(dir, "thumbnails")
		dest := filepath.Join(dir, newFileName)

		// create folder
		os.MkdirAll(dir, os.ModePerm)
		os.MkdirAll(dirThumbnails, os.ModePerm)

		// save picture from form to the dest folder
		if err := c.SaveUploadedFile(file, dest); err != nil {
			t.BadRequest(c, fmt.Sprintf("get form err: %s", fmt.Sprintf("upload file err: %s", err.Error())))
			return
		}

		// open the image and then create the thumbnail
		src, _ := imaging.Open(dest)
		dstImage1920 := imaging.Resize(src, 1920, 0, imaging.Lanczos)
		largePath := filepath.Join(dirThumbnails, "large_"+newFileName)
		dstImage800 := imaging.Resize(src, 800, 0, imaging.Lanczos)
		mediumPath := filepath.Join(dirThumbnails, "medium_"+newFileName)
		dstImage10 := imaging.Resize(src, 10, 10, imaging.Lanczos)
		thumbPath := filepath.Join(dirThumbnails, "thumb_"+newFileName)

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

		// save the data to db

		// response, err := fileService.CreateImage(c, &file.CreateImageRequest{
		// 	New: &file.Image{
		// 		Title: createRoleRequest.Title,
		// 	},
		// })

		// if err != nil {
		// 	t.BadRequest(c, err.Error())
		// 	return
		// }

		// var picture models.Picture
		// picture.Filename = newFileName
		// picture.Path = dest
		// picture.InfoName = file.Filename
		// picture.InfoType = fileType
		// picture.InfoSize = file.Size
		// picture.Large = largePath
		// picture.Medium = mediumPath
		// picture.Thumb = thumbPath

		// db.Create(&picture)
		// pictures = append(pictures, picture)

	}
	t.OKSingleData(c, nil)

	return
}
