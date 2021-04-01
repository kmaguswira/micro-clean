package controllers

import (
	"github.com/asim/go-micro/v3/client"
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
	"github.com/kmaguswira/micro-clean/utils"
)

type emailTemplateController struct {
	utils.Response
	notificationService notification.NotificationService
}

func NewEmailTemplateController(client client.Client) emailTemplateController {
	return emailTemplateController{
		notificationService: notification.NewNotificationService("kmaguswira.srv.notification", client),
	}
}

func (t *emailTemplateController) Create(c *gin.Context) {
	var createEmailTemplateRequest requests.CreateEmailTemplate

	if err := c.BindJSON(&createEmailTemplateRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.notificationService.CreateEmailTemplate(c, &notification.CreateEmailTemplateRequest{
		New: &notification.EmailTemplate{
			Title:     createEmailTemplateRequest.Title,
			Subject:   createEmailTemplateRequest.Subject,
			HTML:      createEmailTemplateRequest.HTML,
			PlainText: createEmailTemplateRequest.PlainText,
			Language:  createEmailTemplateRequest.Language,
			FromName:  createEmailTemplateRequest.FromName,
			FromEmail: createEmailTemplateRequest.FromEmail,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t *emailTemplateController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := t.notificationService.FindEmailTemplateById(c, &notification.FindEmailTemplateByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *emailTemplateController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := t.notificationService.DeleteEmailTemplate(c, &notification.DeleteEmailTemplateRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *emailTemplateController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateEmailTemplateRequest requests.UpdateEmailTemplate

	if err := c.BindJSON(&updateEmailTemplateRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.notificationService.UpdateEmailTemplate(c, &notification.UpdateEmailTemplateRequest{
		Update: &notification.EmailTemplate{
			ID:        id,
			Title:     updateEmailTemplateRequest.Title,
			Subject:   updateEmailTemplateRequest.Subject,
			HTML:      updateEmailTemplateRequest.HTML,
			PlainText: updateEmailTemplateRequest.PlainText,
			Language:  updateEmailTemplateRequest.Language,
			FromName:  updateEmailTemplateRequest.FromName,
			FromEmail: updateEmailTemplateRequest.FromEmail,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *emailTemplateController) FindAll(c *gin.Context) {
	q := utils.QueryBuilder(c)

	response, err := t.notificationService.FindAllEmailTemplate(c, &notification.FindAllEmailTemplateRequest{
		Query: &notification.BaseQuery{
			QueryKey:   q.QueryKey,
			QueryValue: q.QueryValue,
			Limit:      q.Limit,
			Offset:     q.Offset,
			Sort:       q.Sort,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *emailTemplateController) SendEmail(c *gin.Context) {
	var sendEmailRequest requests.SendEmail

	if err := c.BindJSON(&sendEmailRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.notificationService.SendEmail(c, &notification.SendEmailRequest{
		TemplateTitle: sendEmailRequest.TemplateTitle,
		ToName:        sendEmailRequest.ToName,
		ToEmail:       sendEmailRequest.ToEmail,
		Data:          sendEmailRequest.Data,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}
