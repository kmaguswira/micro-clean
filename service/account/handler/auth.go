package handler

import (
	"context"

	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/application/usecases"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/micro/go-micro/util/log"
)

func (t *Account) SignUp(ctx context.Context, req *account.SignUpRequest, res *account.SignUpResponse) error {
	log.Log("Received Account.SignUp request")

	userRole, err := t.findRoleByTitleUseCase.Execute(global.USER_ROLE)

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
		ID:   result.User.ID,
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

	userResponse := populateUserResponse(result.User)
	res.ResponseInfo = t.response.OK()
	res.Result = &userResponse
	return nil
}

func (t *Account) ActivateUser(ctx context.Context, req *account.ActivateUserRequest, res *account.ActivateUserResponse) error {
	log.Log("Received Account.ActivateUser")

	_, err := t.activateUserUseCase.Execute(req.Token)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	res.ResponseInfo = t.response.OK()

	return nil
}

func (t *Account) ForgotPassword(ctx context.Context, req *account.ForgotPasswordRequest, res *account.ForgotPasswordResponse) error {
	log.Log("Received Account.ForgotPassword")

	_, err := t.forgotPasswordUseCase.Execute(req.Email)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	//TODO: send email

	res.ResponseInfo = t.response.OK()
	return nil
}

func (t *Account) ResetPassword(ctx context.Context, req *account.ResetPasswordRequest, res *account.ResetPasswordResponse) error {
	log.Log("Received Account.ResetPassword")

	_, err := t.resetPasswordUseCase.Execute(req.Token, req.Password)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	res.ResponseInfo = t.response.OK()
	return nil
}

func (t *Account) ChangePassword(ctx context.Context, req *account.ChangePasswordRequest, res *account.ChangePasswordResponse) error {
	log.Log("Received Account.ChangePassword")

	_, err := t.changePasswordUseCase.Execute(req.UserID, req.OldPassword, req.NewPassword)

	if err != nil {
		res.ResponseInfo = t.response.InternalServerError()
		return nil
	}

	res.ResponseInfo = t.response.OK()
	return nil
}
