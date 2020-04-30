package utils

import (
	_proto "github.com/kmaguswira/micro-clean/service/file/proto/file"
)

type Response struct{}

func (r Response) OK() *_proto.BaseResponse {
	message := _proto.BaseResponse{
		Status:  "200",
		Message: "Ok",
	}
	return &message
}

func (r Response) Created() *_proto.BaseResponse {
	message := _proto.BaseResponse{
		Status:  "201",
		Message: "Created",
	}
	return &message
}

func (r Response) BadRequest() *_proto.BaseResponse {
	message := _proto.BaseResponse{
		Status:  "400",
		Message: "Bad Request",
	}
	return &message
}

func (r Response) InternalServerError() *_proto.BaseResponse {
	message := _proto.BaseResponse{
		Status:  "500",
		Message: "Internal Server Error",
	}
	return &message
}
