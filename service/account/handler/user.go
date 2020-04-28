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

	userResponse := populateUserResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &userResponse

	return nil
}

func (t *Account) FindUserById(ctx context.Context, req *account.FindUserByIdRequest, res *account.FindUserByIdResponse) error {
	log.Log("Received Account.FindUserById")

	result, err := t.findUserByIDUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	userResponse := populateUserResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &userResponse

	return nil
}

func (t *Account) FindAllUser(ctx context.Context, req *account.FindAllUserRequest, res *account.FindAllUserResponse) error {
	log.Log("Received Account.FindAllUser")

	input := global.FindAllInput{
		QueryKey: req.Query.QueryKey,
		Limit:    req.Query.Limit,
		Offset:   req.Query.Offset,
		Sort:     req.Query.Sort,
	}

	input.ParseValue(req.Query.QueryValue)

	result, err := t.findAllUserUseCase.Execute(input)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	var users []*account.User
	From(result).Select(func(c interface{}) interface{} {
		d := c.(domain.User)
		return populateUserResponse(d)
	}).ToSlice(&users)

	res.ResponseInfo = t.response.OK()
	res.Result = users
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

	userResponse := populateUserResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &userResponse

	return nil
}

func (t *Account) DeleteUser(ctx context.Context, req *account.DeleteUserRequest, res *account.DeleteUserResponse) error {
	log.Log("Received Account.DeleteUser")

	result, err := t.deleteUserUseCase.Execute(req.Id)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	userResponse := populateUserResponse(result)
	res.ResponseInfo = t.response.OK()
	res.Result = &userResponse

	return nil
}

func (t *Account) ActivateUser(ctx context.Context, req *account.ActivateUserRequest, res *account.ActivateUserResponse) error {
	log.Log("Received Account.DeleteUser")

	_, err := t.activateUserUseCase.Execute(req.Token)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	res.ResponseInfo = t.response.OK()

	return nil
}

func (t *Account) ForgotPassword(ctx context.Context, req *account.ForgotPasswordRequest, res *account.ForgotPasswordResponse) error {
	return nil
}

func (t *Account) ResetPassword(ctx context.Context, req *account.ResetPasswordRequest, res *account.ResetPasswordResponse) error {
	return nil
}

func (t *Account) ChangePassword(ctx context.Context, req *account.ChangePasswordRequest, res *account.ChangePasswordResponse) error {
	return nil
}
