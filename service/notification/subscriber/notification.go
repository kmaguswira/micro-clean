package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
)

type Notification struct{}

func (e *Notification) Handle(ctx context.Context, msg *notification.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *notification.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
