package handler

import (
	"context"

	"github.com/kmaguswira/micro-clean/service/account/application/usecases"
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
		ID:    result.ID,
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
		ID:    result.ID,
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
	log.Log("Received Account.UpdateRole")

	input := usecases.UpdateRoleInput{
		ID:    req.Update.ID,
		Title: req.Update.Title,
	}

	result, err := t.updateRoleUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.Role{
		ID:    result.ID,
		Title: result.Title,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}

func (t *Account) DeleteRole(ctx context.Context, req *account.DeleteRoleRequest, res *account.DeleteRoleResponse) error {
	log.Log("Received Account.DeleteRole")

	result, err := t.deleteRoleUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.Role{
		ID:    result.ID,
		Title: result.Title,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

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
		ID:    result.ID,
		Title: result.Title,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}
