package handler

import (
	"context"

	"github.com/kmaguswira/micro-clean/service/account/application/usecases"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/micro/go-micro/util/log"
)

func (t *Account) CreateUser(ctx context.Context, req *account.CreateUserRequest, res *account.CreateUserResponse) error {
	log.Log("Received Account.CreateUser")

	input := usecases.CreateUserInput{
		Name:     req.New.Name,
		Email:    req.New.Email,
		Username: req.New.Username,
		Password: req.New.Password,
		Status:   req.New.Status,
		RoleID:   req.New.RoleID,
	}

	result, err := t.createUserUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.User{
		ID:       result.ID,
		Name:     result.Name,
		Email:    result.Email,
		Username: result.Username,
		Status:   result.Status,
		RoleID:   result.RoleID,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}

func (t *Account) FindUserById(ctx context.Context, req *account.FindUserByIdRequest, res *account.FindUserByIdResponse) error {
	log.Log("Received Account.FindUserById")

	result, err := t.findUserByIDUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.User{
		ID:       result.ID,
		Name:     result.Name,
		Email:    result.Email,
		Username: result.Username,
		Status:   result.Status,
		RoleID:   result.RoleID,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}

func (t *Account) FindAllUser(ctx context.Context, req *account.FindAllUserRequest, res *account.FindAllUserResponse) error {
	return nil
}

func (t *Account) UpdateUser(ctx context.Context, req *account.UpdateUserRequest, res *account.UpdateUserResponse) error {
	log.Log("Received Account.UpdateUser")

	input := usecases.UpdateUserInput{
		ID:       req.Update.ID,
		Name:     req.Update.Name,
		Username: req.Update.Username,
		Email:    req.Update.Email,
		RoleID:   req.Update.RoleID,
		Status:   req.Update.Status,
	}

	result, err := t.updateUserUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.User{
		ID:       result.ID,
		Name:     result.Name,
		Email:    result.Email,
		Username: result.Username,
		Status:   result.Status,
		RoleID:   result.RoleID,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}

func (t *Account) DeleteUser(ctx context.Context, req *account.DeleteUserRequest, res *account.DeleteUserResponse) error {
	log.Log("Received Account.DeleteUser")

	result, err := t.deleteUserUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	Result := account.User{
		ID:       result.ID,
		Name:     result.Name,
		Email:    result.Email,
		Username: result.Username,
		Status:   result.Status,
		RoleID:   result.RoleID,
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}
