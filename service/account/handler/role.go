package handler

import (
	"context"

	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/micro/go-micro/util/log"
)

func (t *Account) CreateRole(ctx context.Context, req *account.CreateRoleRequest, res *account.CreateRoleResponse) error {
	log.Log("Received Account.CreateRole request")
	result, err := t.createRoleUseCase.Execute(req.New.Title)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.Role{
		Id:    result.ID,
		Title: result.Title,
	}
	res.ResponseInfo = t.response.Created()
	res.Result = &Result

	return nil
}

func (t *Account) FindRoleById(ctx context.Context, req *account.FindRoleByIdRequest, res *account.FindRoleByIdResponse) error {
	log.Log("Received Account.FindRoleById")

	result, err := t.findRoleByIDUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.Role{
		Id:    result.ID,
		Title: result.Title,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}

func (t *Account) FindAllRole(ctx context.Context, req *account.FindAllRoleRequest, res *account.FindAllRoleResponse) error {
	return nil
}

func (t *Account) UpdateRole(ctx context.Context, req *account.UpdateRoleRequest, res *account.UpdateRoleResponse) error {
	return nil
}

func (t *Account) DeleteRole(ctx context.Context, req *account.DeleteRoleRequest, res *account.DeleteRoleResponse) error {
	return nil
}

func (t *Account) FindRoleByTitle(ctx context.Context, req *account.FindRoleByTitleRequest, res *account.FindRoleByTitleResponse) error {
	log.Log("Received Account.FindRoleByTitle request")

	result, err := t.findRoleByTitleUseCase.Execute(req.Title)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.Role{
		Id:    result.ID,
		Title: result.Title,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}
