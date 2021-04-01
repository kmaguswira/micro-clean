package handler

import (
	"context"

	"github.com/asim/go-micro/v3/util/log"
	"github.com/kmaguswira/micro-clean/service/file/application/usecases"
	"github.com/kmaguswira/micro-clean/service/file/domain"
	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
	"github.com/kmaguswira/micro-clean/service/file/repositories"
	"github.com/kmaguswira/micro-clean/service/file/utils"
)

type File struct {
	response                utils.Response
	createImageUseCase      usecases.ICreateImage
	findImageByIDUseCase    usecases.IFindImageByID
	deleteImageUseCase      usecases.IDeleteImage
	updateImageUseCase      usecases.IUpdateImage
	findAllImageUseCase     usecases.IFindAllImage
	createDocumentUseCase   usecases.ICreateDocument
	findDocumentByIDUseCase usecases.IFindDocumentByID
	deleteDocumentUseCase   usecases.IDeleteDocument
	updateDocumentUseCase   usecases.IUpdateDocument
	findAllDocumentUseCase  usecases.IFindAllDocument
}

func NewFile() *File {
	readWriteRepository := repositories.NewReadWriteRepository(nil)
	readRepository := repositories.NewReadRepository(nil)

	return &File{
		createImageUseCase:      usecases.NewCreateImageUseCase(readWriteRepository),
		findImageByIDUseCase:    usecases.NewFindImageByIDUseCase(readRepository),
		deleteImageUseCase:      usecases.NewDeleteImageUseCase(readWriteRepository),
		updateImageUseCase:      usecases.NewUpdateImageUseCase(readWriteRepository),
		findAllImageUseCase:     usecases.NewFindAllImageUseCase(readRepository),
		createDocumentUseCase:   usecases.NewCreateDocumentUseCase(readWriteRepository),
		findDocumentByIDUseCase: usecases.NewFindDocumentByIDUseCase(readRepository),
		deleteDocumentUseCase:   usecases.NewDeleteDocumentUseCase(readWriteRepository),
		updateDocumentUseCase:   usecases.NewUpdateDocumentUseCase(readWriteRepository),
		findAllDocumentUseCase:  usecases.NewFindAllDocumentUseCase(readRepository),
		response:                utils.Response{},
	}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *File) Call(ctx context.Context, req *file.Request, rsp *file.Response) error {
	log.Log("Received File.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *File) Stream(ctx context.Context, req *file.StreamingRequest, stream file.File_StreamStream) error {
	log.Logf("Received File.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&file.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *File) PingPong(ctx context.Context, stream file.File_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&file.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func populateImageResponse(image domain.Image) file.Image {
	imageResponse := file.Image{
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

	return imageResponse
}

func populateDocumentResponse(document domain.Document) file.Document {
	documentResponse := file.Document{
		ID:          document.ID,
		Name:        document.Name,
		Path:        document.Path,
		Slug:        document.Slug,
		Type:        document.Type,
		Description: document.Description,
		Info:        document.Info,
		Cdn:         document.Cdn,
	}

	return documentResponse
}
