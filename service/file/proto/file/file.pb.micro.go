// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/file/file.proto

package go_micro_srv_file

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for File service

type FileService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (File_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (File_PingPongService, error)
	CreateImage(ctx context.Context, in *CreateImageRequest, opts ...client.CallOption) (*CreateImageResponse, error)
	FindImageById(ctx context.Context, in *FindImageByIdRequest, opts ...client.CallOption) (*FindImageByIdResponse, error)
	FindAllImage(ctx context.Context, in *FindAllImageRequest, opts ...client.CallOption) (*FindAllImageResponse, error)
	UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...client.CallOption) (*UpdateImageResponse, error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...client.CallOption) (*DeleteImageResponse, error)
	CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...client.CallOption) (*CreateDocumentResponse, error)
	FindDocumentById(ctx context.Context, in *FindDocumentByIdRequest, opts ...client.CallOption) (*FindDocumentByIdResponse, error)
	FindAllDocument(ctx context.Context, in *FindAllDocumentRequest, opts ...client.CallOption) (*FindAllDocumentResponse, error)
	UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...client.CallOption) (*UpdateDocumentResponse, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...client.CallOption) (*DeleteDocumentResponse, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.file"
	}
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "File.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (File_StreamService, error) {
	req := c.c.NewRequest(c.name, "File.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &fileServiceStream{stream}, nil
}

type File_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type fileServiceStream struct {
	stream client.Stream
}

func (x *fileServiceStream) Close() error {
	return x.stream.Close()
}

func (x *fileServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileService) PingPong(ctx context.Context, opts ...client.CallOption) (File_PingPongService, error) {
	req := c.c.NewRequest(c.name, "File.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &fileServicePingPong{stream}, nil
}

type File_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type fileServicePingPong struct {
	stream client.Stream
}

func (x *fileServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *fileServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *fileServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileService) CreateImage(ctx context.Context, in *CreateImageRequest, opts ...client.CallOption) (*CreateImageResponse, error) {
	req := c.c.NewRequest(c.name, "File.CreateImage", in)
	out := new(CreateImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) FindImageById(ctx context.Context, in *FindImageByIdRequest, opts ...client.CallOption) (*FindImageByIdResponse, error) {
	req := c.c.NewRequest(c.name, "File.FindImageById", in)
	out := new(FindImageByIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) FindAllImage(ctx context.Context, in *FindAllImageRequest, opts ...client.CallOption) (*FindAllImageResponse, error) {
	req := c.c.NewRequest(c.name, "File.FindAllImage", in)
	out := new(FindAllImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...client.CallOption) (*UpdateImageResponse, error) {
	req := c.c.NewRequest(c.name, "File.UpdateImage", in)
	out := new(UpdateImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...client.CallOption) (*DeleteImageResponse, error) {
	req := c.c.NewRequest(c.name, "File.DeleteImage", in)
	out := new(DeleteImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...client.CallOption) (*CreateDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "File.CreateDocument", in)
	out := new(CreateDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) FindDocumentById(ctx context.Context, in *FindDocumentByIdRequest, opts ...client.CallOption) (*FindDocumentByIdResponse, error) {
	req := c.c.NewRequest(c.name, "File.FindDocumentById", in)
	out := new(FindDocumentByIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) FindAllDocument(ctx context.Context, in *FindAllDocumentRequest, opts ...client.CallOption) (*FindAllDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "File.FindAllDocument", in)
	out := new(FindAllDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...client.CallOption) (*UpdateDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "File.UpdateDocument", in)
	out := new(UpdateDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...client.CallOption) (*DeleteDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "File.DeleteDocument", in)
	out := new(DeleteDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for File service

type FileHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, File_StreamStream) error
	PingPong(context.Context, File_PingPongStream) error
	CreateImage(context.Context, *CreateImageRequest, *CreateImageResponse) error
	FindImageById(context.Context, *FindImageByIdRequest, *FindImageByIdResponse) error
	FindAllImage(context.Context, *FindAllImageRequest, *FindAllImageResponse) error
	UpdateImage(context.Context, *UpdateImageRequest, *UpdateImageResponse) error
	DeleteImage(context.Context, *DeleteImageRequest, *DeleteImageResponse) error
	CreateDocument(context.Context, *CreateDocumentRequest, *CreateDocumentResponse) error
	FindDocumentById(context.Context, *FindDocumentByIdRequest, *FindDocumentByIdResponse) error
	FindAllDocument(context.Context, *FindAllDocumentRequest, *FindAllDocumentResponse) error
	UpdateDocument(context.Context, *UpdateDocumentRequest, *UpdateDocumentResponse) error
	DeleteDocument(context.Context, *DeleteDocumentRequest, *DeleteDocumentResponse) error
}

func RegisterFileHandler(s server.Server, hdlr FileHandler, opts ...server.HandlerOption) error {
	type file interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
		CreateImage(ctx context.Context, in *CreateImageRequest, out *CreateImageResponse) error
		FindImageById(ctx context.Context, in *FindImageByIdRequest, out *FindImageByIdResponse) error
		FindAllImage(ctx context.Context, in *FindAllImageRequest, out *FindAllImageResponse) error
		UpdateImage(ctx context.Context, in *UpdateImageRequest, out *UpdateImageResponse) error
		DeleteImage(ctx context.Context, in *DeleteImageRequest, out *DeleteImageResponse) error
		CreateDocument(ctx context.Context, in *CreateDocumentRequest, out *CreateDocumentResponse) error
		FindDocumentById(ctx context.Context, in *FindDocumentByIdRequest, out *FindDocumentByIdResponse) error
		FindAllDocument(ctx context.Context, in *FindAllDocumentRequest, out *FindAllDocumentResponse) error
		UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, out *UpdateDocumentResponse) error
		DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, out *DeleteDocumentResponse) error
	}
	type File struct {
		file
	}
	h := &fileHandler{hdlr}
	return s.Handle(s.NewHandler(&File{h}, opts...))
}

type fileHandler struct {
	FileHandler
}

func (h *fileHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.FileHandler.Call(ctx, in, out)
}

func (h *fileHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FileHandler.Stream(ctx, m, &fileStreamStream{stream})
}

type File_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type fileStreamStream struct {
	stream server.Stream
}

func (x *fileStreamStream) Close() error {
	return x.stream.Close()
}

func (x *fileStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *fileHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.FileHandler.PingPong(ctx, &filePingPongStream{stream})
}

type File_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type filePingPongStream struct {
	stream server.Stream
}

func (x *filePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *filePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *filePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *filePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *filePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *fileHandler) CreateImage(ctx context.Context, in *CreateImageRequest, out *CreateImageResponse) error {
	return h.FileHandler.CreateImage(ctx, in, out)
}

func (h *fileHandler) FindImageById(ctx context.Context, in *FindImageByIdRequest, out *FindImageByIdResponse) error {
	return h.FileHandler.FindImageById(ctx, in, out)
}

func (h *fileHandler) FindAllImage(ctx context.Context, in *FindAllImageRequest, out *FindAllImageResponse) error {
	return h.FileHandler.FindAllImage(ctx, in, out)
}

func (h *fileHandler) UpdateImage(ctx context.Context, in *UpdateImageRequest, out *UpdateImageResponse) error {
	return h.FileHandler.UpdateImage(ctx, in, out)
}

func (h *fileHandler) DeleteImage(ctx context.Context, in *DeleteImageRequest, out *DeleteImageResponse) error {
	return h.FileHandler.DeleteImage(ctx, in, out)
}

func (h *fileHandler) CreateDocument(ctx context.Context, in *CreateDocumentRequest, out *CreateDocumentResponse) error {
	return h.FileHandler.CreateDocument(ctx, in, out)
}

func (h *fileHandler) FindDocumentById(ctx context.Context, in *FindDocumentByIdRequest, out *FindDocumentByIdResponse) error {
	return h.FileHandler.FindDocumentById(ctx, in, out)
}

func (h *fileHandler) FindAllDocument(ctx context.Context, in *FindAllDocumentRequest, out *FindAllDocumentResponse) error {
	return h.FileHandler.FindAllDocument(ctx, in, out)
}

func (h *fileHandler) UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, out *UpdateDocumentResponse) error {
	return h.FileHandler.UpdateDocument(ctx, in, out)
}

func (h *fileHandler) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, out *DeleteDocumentResponse) error {
	return h.FileHandler.DeleteDocument(ctx, in, out)
}
