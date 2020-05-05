package handler

import (
	"context"

	"github.com/kmaguswira/micro-clean/service/notification/application/usecases"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
	"github.com/kmaguswira/micro-clean/service/notification/external_service"
	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
	"github.com/kmaguswira/micro-clean/service/notification/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/utils"
	"github.com/micro/go-micro/util/log"
)

type Notification struct {
	createEmailTemplateUseCase   usecases.ICreateEmailTemplate
	findEmailTemplateByIDUseCase usecases.IFindEmailTemplateByID
	deleteEmailTemplateUseCase   usecases.IDeleteEmailTemplate
	updateEmailTemplateUseCase   usecases.IUpdateEmailTemplate
	findAllEmailTemplateUseCase  usecases.IFindAllEmailTemplate
	sendEmailUseCase             usecases.ISendEmail
	response                     utils.Response
}

func NewNotification() *Notification {
	readWriteRepository := repositories.NewReadWriteRepository(nil)
	readRepository := repositories.NewReadRepository(nil)
	sendEmailService := external_service.NewSendGridService()

	return &Notification{
		createEmailTemplateUseCase:   usecases.NewCreateEmailTemplateUseCase(readWriteRepository),
		findEmailTemplateByIDUseCase: usecases.NewFindEmailTemplateByIDUseCase(readRepository),
		deleteEmailTemplateUseCase:   usecases.NewDeleteEmailTemplateUseCase(readWriteRepository),
		updateEmailTemplateUseCase:   usecases.NewUpdateEmailTemplateUseCase(readWriteRepository),
		findAllEmailTemplateUseCase:  usecases.NewFindAllEmailTemplateUseCase(readRepository),
		sendEmailUseCase:             usecases.NewSendEmailUseCase(readRepository, sendEmailService),
		response:                     utils.Response{},
	}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Notification) Call(ctx context.Context, req *notification.Request, rsp *notification.Response) error {
	log.Log("Received Notification.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Notification) Stream(ctx context.Context, req *notification.StreamingRequest, stream notification.Notification_StreamStream) error {
	log.Logf("Received Notification.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&notification.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Notification) PingPong(ctx context.Context, stream notification.Notification_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&notification.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func populateEmailTemplateResponse(emailTemplate domain.EmailTemplate) notification.EmailTemplate {
	emailTemplateResponse := notification.EmailTemplate{
		ID:        emailTemplate.ID,
		Title:     emailTemplate.Title,
		Subject:   emailTemplate.Subject,
		HTML:      emailTemplate.HTML,
		PlainText: emailTemplate.PlainText,
		Language:  emailTemplate.Language,
		FromName:  emailTemplate.FromName,
		FromEmail: emailTemplate.FromEmail,
	}

	return emailTemplateResponse
}
