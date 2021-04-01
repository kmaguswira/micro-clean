package subscriber

import (
	"context"
	"github.com/asim/go-micro/v3/util/log"

	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
)

type Account struct{}

func (e *Account) Handle(ctx context.Context, msg *account.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *account.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
