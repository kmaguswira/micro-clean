package subscriber

import (
	"context"

	"github.com/asim/go-micro/v3/util/log"

	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
)

type File struct{}

func (e *File) Handle(ctx context.Context, msg *file.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *file.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
