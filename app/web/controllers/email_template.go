package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro/client"
)

var notificationService = notification.NewNotificationService("kmaguswira.srv.notification", client.DefaultClient)

type EmailTemplateController struct {
	utils.Response
}

func (t EmailTemplateController) Create(c *gin.Context) {
	var createEmailTemplateRequest requests.CreateEmailTemplate

	if err := c.BindJSON(&createEmailTemplateRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := notificationService.CreateEmailTemplate(c, &notification.CreateEmailTemplateRequest{
		New: &notification.EmailTemplate{
			Title:    createEmailTemplateRequest.Title,
			Subject:  createEmailTemplateRequest.Subject,
			Body:     createEmailTemplateRequest.Body,
			Language: createEmailTemplateRequest.Language,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t EmailTemplateController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := notificationService.FindEmailTemplateById(c, &notification.FindEmailTemplateByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t EmailTemplateController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := notificationService.DeleteEmailTemplate(c, &notification.DeleteEmailTemplateRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t EmailTemplateController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateEmailTemplateRequest requests.UpdateEmailTemplate

	if err := c.BindJSON(&updateEmailTemplateRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := notificationService.UpdateEmailTemplate(c, &notification.UpdateEmailTemplateRequest{
		Update: &notification.EmailTemplate{
			ID:       id,
			Title:    updateEmailTemplateRequest.Title,
			Subject:  updateEmailTemplateRequest.Subject,
			Body:     updateEmailTemplateRequest.Body,
			Language: updateEmailTemplateRequest.Language,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t EmailTemplateController) FindAll(c *gin.Context) {
	q := utils.QueryBuilder(c)

	response, err := notificationService.FindAllEmailTemplate(c, &notification.FindAllEmailTemplateRequest{
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
