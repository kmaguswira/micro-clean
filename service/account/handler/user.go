package handler

import (
	"context"

	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
)

func (t *Account) CreateUser(ctx context.Context, req *account.CreateUserRequest, res *account.CreateUserResponse) error {
	return nil
}

func (t *Account) FindUserById(ctx context.Context, req *account.FindUserByIdRequest, res *account.FindUserByIdResponse) error {
	return nil
}

func (t *Account) FindAllUser(ctx context.Context, req *account.FindAllUserRequest, res *account.FindAllUserResponse) error {
	return nil
}

func (t *Account) UpdateUser(ctx context.Context, req *account.UpdateUserRequest, res *account.UpdateUserResponse) error {
	return nil
}

func (t *Account) DeleteUser(ctx context.Context, req *account.DeleteUserRequest, res *account.DeleteUserResponse) error {
	return nil
}
