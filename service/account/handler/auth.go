package handler

import (
	"context"

	"github.com/kmaguswira/micro-clean/service/account/application/usecases"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/micro/go-micro/util/log"
)

func (t *Account) SignUp(ctx context.Context, req *account.SignUpRequest, res *account.SignUpResponse) error {
	log.Log("Received Account.SignUp request")

	// change to constant
	userRole, err := t.findRoleByTitleUseCase.Execute("user")

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	newUser := usecases.SignUpInput{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
		Role:     userRole.ID,
		Name:     req.Name,
	}

	result, err := t.signUpUseCase.Execute(newUser)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	res.ResponseInfo = t.response.Created()
	res.Result = &account.User{
		Id:   result.User.ID,
		Name: result.User.Name,
	}
	return nil
}

func (t *Account) SignIn(ctx context.Context, req *account.SignInRequest, res *account.SignInResponse) error {
	log.Log("Received Account.SignIn request")

	newSiginInInput := usecases.SignInInput{
		User:     req.User,
		Password: req.Password,
	}

	result, err := t.signInUseCase.Execute(newSiginInInput)

	if err != nil {
		res.ResponseInfo = t.response.BadRequest()
		return nil
	}

	res.ResponseInfo = t.response.OK()
	res.Result = &account.User{
		Id:   result.User.ID,
		Name: result.User.Name,
	}
	return nil
}
