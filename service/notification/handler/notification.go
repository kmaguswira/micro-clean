package handler

import (
	"context"

	. "github.com/ahmetb/go-linq"
	"github.com/kmaguswira/micro-clean/service/notification/application/global"
	"github.com/kmaguswira/micro-clean/service/notification/application/usecases"
	"github.com/kmaguswira/micro-clean/service/notification/domain"
	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
	"github.com/micro/go-micro/util/log"
)

func (t *Notification) CreateEmailTemplate(ctx context.Context, req *notification.CreateEmailTemplateRequest, res *notification.CreateEmailTemplateResponse) error {
	log.Log("Received Notification.CreateEmailTemplate request")

	input := usecases.CreateEmailTemplateInput{
		Title:     req.New.Title,
		Subject:   req.New.Subject,
		HTML:      req.New.HTML,
		PlainText: req.New.PlainText,
		Language:  req.New.Language,
		FromName:  req.New.FromName,
		FromEmail: req.New.FromEmail,
	}

	result, err := t.createEmailTemplateUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	emailTemplateResponse := populateEmailTemplateResponse(result)
	res.ResponseInfo = t.response.Created()
	res.Result = &emailTemplateResponse

	return nil
}

func (t *Notification) FindEmailTemplateById(ctx context.Context, req *notification.FindEmailTemplateByIdRequest, res *notification.FindEmailTemplateByIdResponse) error {
	log.Log("Received Notification.FindEmailTemplateById")

	result, err := t.findEmailTemplateByIDUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	emailTemplateResponse := populateEmailTemplateResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &emailTemplateResponse

	return nil
}

func (t *Notification) FindAllEmailTemplate(ctx context.Context, req *notification.FindAllEmailTemplateRequest, res *notification.FindAllEmailTemplateResponse) error {
	log.Log("Received Notification.FindAllEmailTemplate")

	input := global.FindAllInput{
		QueryKey: req.Query.QueryKey,
		Limit:    req.Query.Limit,
		Offset:   req.Query.Offset,
		Sort:     req.Query.Sort,
	}

	input.ParseValue(req.Query.QueryValue)

	result, err := t.findAllEmailTemplateUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	var emailTemplates []*notification.EmailTemplate
	From(result).Select(func(c interface{}) interface{} {
		d := c.(domain.EmailTemplate)
		emailTemplate := populateEmailTemplateResponse(d)
		return &emailTemplate
	}).ToSlice(&emailTemplates)

	res.ResponseInfo = t.response.OK()
	res.Result = emailTemplates
	return nil
}

func (t *Notification) UpdateEmailTemplate(ctx context.Context, req *notification.UpdateEmailTemplateRequest, res *notification.UpdateEmailTemplateResponse) error {
	log.Log("Received Notification.UpdateEmailTemplate")

	input := usecases.UpdateEmailTemplateInput{
		ID:        req.Update.ID,
		Title:     req.Update.Title,
		Subject:   req.Update.Subject,
		HTML:      req.Update.HTML,
		PlainText: req.Update.PlainText,
		Language:  req.Update.Language,
		FromName:  req.Update.FromName,
		FromEmail: req.Update.FromEmail,
	}

	result, err := t.updateEmailTemplateUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	emailTemplateResponse := populateEmailTemplateResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &emailTemplateResponse

	return nil
}

func (t *Notification) DeleteEmailTemplate(ctx context.Context, req *notification.DeleteEmailTemplateRequest, res *notification.DeleteEmailTemplateResponse) error {
	log.Log("Received Notification.DeleteEmailTemplate")

	result, err := t.deleteEmailTemplateUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	emailTemplateResponse := populateEmailTemplateResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &emailTemplateResponse

	return nil
}

func (t *Notification) SendEmail(ctx context.Context, req *notification.SendEmailRequest, res *notification.SendEmailResponse) error {
	log.Log("Received Notification.SendEmail")

	data := []interface{}{}
	for i := 0; i < len(req.Data); i++ {
		data = append(data, req.Data[i])
	}

	_, err := t.sendEmailUseCase.Execute(req.TemplateTitle, req.ToName, req.ToEmail, data)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	res.ResponseInfo = t.response.OK()
	return nil
}
