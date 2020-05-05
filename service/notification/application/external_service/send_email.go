package external_service

import (
	"github.com/kmaguswira/micro-clean/service/notification/application/global"
)

type ISendEmail interface {
	SendEmail(input global.SendEmailInput) (bool, error)
}
