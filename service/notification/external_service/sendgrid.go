package external_service

import (
	"log"

	iface "github.com/kmaguswira/micro-clean/service/notification/application/external_service"
	"github.com/kmaguswira/micro-clean/service/notification/application/global"
	"github.com/kmaguswira/micro-clean/service/notification/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridService struct {
	client *sendgrid.Client
}

func NewSendGridService() iface.ISendEmail {
	cfg := config.GetConfig()
	return &SendGridService{
		client: sendgrid.NewSendClient(cfg.ExternalServices.Sendgrid.APIKEY),
	}
}

func (t *SendGridService) SendEmail(input global.SendEmailInput) (bool, error) {
	from := mail.NewEmail(input.From.Name, input.From.Email)
	to := mail.NewEmail(input.To.Name, input.To.Email)
	message := mail.NewSingleEmail(from, input.Subject, to, input.PlainTextContent, input.HTML)
	response, err := t.client.Send(message)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println(response.StatusCode)

	return true, nil
}
