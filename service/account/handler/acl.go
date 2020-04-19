package handler

import (
	"context"
	"strings"

	. "github.com/ahmetb/go-linq"
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

	var roleID []string
	From(result.Permitted).Select(func(c interface{}) interface{} {
		return c.(domain.Role).ID
	}).ToSlice(&roleID)

	Result := account.ACL{
		ID:        result.ID,
		Title:     result.Title,
		Handler:   result.Handler,
		IsPublic:  result.IsPublic,
		Permitted: strings.Join(roleID, ","),
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &Result

	return nil
}

func (t *Account) FindACLById(ctx context.Context, req *account.FindACLByIdRequest, res *account.FindACLByIdResponse) error {
	return nil
}

func (t *Account) FindAllACL(ctx context.Context, req *account.FindAllACLRequest, res *account.FindAllACLResponse) error {
	return nil
}

func (t *Account) UpdateACL(ctx context.Context, req *account.UpdateACLRequest, res *account.UpdateACLResponse) error {
	return nil
}

func (t *Account) DeleteACL(ctx context.Context, req *account.DeleteACLRequest, res *account.DeleteACLResponse) error {
	return nil
}
