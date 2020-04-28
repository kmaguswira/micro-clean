package handler

import (
	"context"

	. "github.com/ahmetb/go-linq"
	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/application/usecases"
	"github.com/kmaguswira/micro-clean/service/account/domain"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/micro/go-micro/util/log"
)

func (t *Account) CreateACL(ctx context.Context, req *account.CreateACLRequest, res *account.CreateACLResponse) error {
	log.Log("Received Account.CreateACL")

	input := usecases.CreateACLInput{
		Title:     req.New.Title,
		Handler:   req.New.Handler,
		IsPublic:  req.New.IsPublic,
		Permitted: req.New.Permitted,
	}

	result, err := t.createACLUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	aclResponse := populateACLResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &aclResponse

	return nil
}

func (t *Account) FindACLById(ctx context.Context, req *account.FindACLByIdRequest, res *account.FindACLByIdResponse) error {
	log.Log("Received Account.FindACLById")

	result, err := t.findACLByIDUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	aclResponse := populateACLResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &aclResponse

	return nil
}

func (t *Account) FindAllACL(ctx context.Context, req *account.FindAllACLRequest, res *account.FindAllACLResponse) error {
	log.Log("Received Account.FindAllACL")

	input := global.FindAllInput{
		QueryKey: req.Query.QueryKey,
		Limit:    req.Query.Limit,
		Offset:   req.Query.Offset,
		Sort:     req.Query.Sort,
	}

	input.ParseValue(req.Query.QueryValue)

	result, err := t.findAllACLUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	var acls []*account.ACL
	From(result).Select(func(c interface{}) interface{} {
		d := c.(domain.ACL)
		acl := populateACLResponse(d)
		return &acl
	}).ToSlice(&acls)

	res.ResponseInfo = t.response.OK()
	res.Result = acls
	return nil
}

func (t *Account) UpdateACL(ctx context.Context, req *account.UpdateACLRequest, res *account.UpdateACLResponse) error {
	log.Log("Received Account.UpdateACL")

	input := usecases.UpdateACLInput{
		ID:        req.Update.ID,
		Title:     req.Update.Title,
		Handler:   req.Update.Handler,
		IsPublic:  req.Update.IsPublic,
		Permitted: req.Update.Permitted,
	}

	result, err := t.updateACLUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	aclResponse := populateACLResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &aclResponse

	return nil
}

func (t *Account) DeleteACL(ctx context.Context, req *account.DeleteACLRequest, res *account.DeleteACLResponse) error {
	log.Log("Received Account.DeleteACL")

	result, err := t.deleteACLUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	aclResponse := populateACLResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &aclResponse

	return nil
}
