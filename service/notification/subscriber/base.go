package subscriber

import (
	"github.com/kmaguswira/micro-clean/service/notification/application/usecases"
	"github.com/kmaguswira/micro-clean/service/notification/external_service"
	"github.com/kmaguswira/micro-clean/service/notification/repositories"
)

type Notification struct {
	sendEmailUseCase usecases.ISendEmail
}

func NewNotification() *Notification {
	readRepository := repositories.NewReadRepository(nil)
	sendEmailService := external_service.NewSendGridService()

	return &Notification{
		sendEmailUseCase: usecases.NewSendEmailUseCase(readRepository, sendEmailService),
	}
}
