package subscriber

import (
	"context"

	"github.com/asim/go-micro/v3/util/log"

	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
)

func (t *Notification) SendEmail(ctx context.Context, req *notification.SendEmailRequest) error {
	log.Log("Handler Received message: ", req.TemplateTitle)
	data := []interface{}{}
	for i := 0; i < len(req.Data); i++ {
		data = append(data, req.Data[i])
	}

	t.sendEmailUseCase.Execute(req.TemplateTitle, req.ToName, req.ToEmail, data)

	return nil
}

func Handler(ctx context.Context, req *notification.SendEmailRequest) error {
	log.Log("Function Received message: ", req.TemplateTitle)
	return nil
}
