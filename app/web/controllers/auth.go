package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/config"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro/client"
)

var accountService = account.NewAccountService("kmaguswira.srv.account", client.DefaultClient)

type AuthController struct {
	utils.Response
}

func (t AuthController) SignIn(c *gin.Context) {
	var signInRequest requests.SignIn

	if err := c.BindJSON(&signInRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := accountService.SignIn(c, &account.SignInRequest{
		User:     signInRequest.User,
		Password: signInRequest.Password,
	})

	if response.ResponseInfo.Status != "200" {
		fmt.Println(response.ResponseInfo.Status)
		t.BadRequest(c, response.ResponseInfo.Message)
		return

	}

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	cfg := config.GetConfig()
	token := utils.GenerateToken(response.Result.ID, response.Result.RoleID, cfg.Server.Secret)

	t.OKSingleData(c, token)
	return
}

func (t AuthController) SignUp(c *gin.Context) {
	var signUpRequest requests.SignUp

	if err := c.BindJSON(&signUpRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := accountService.SignUp(c, &account.SignUpRequest{
		Email:    signUpRequest.Email,
		Username: signUpRequest.Username,
		Password: signUpRequest.Password,
		Name:     signUpRequest.Name,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t AuthController) ActivateUser(c *gin.Context) {
	token := c.Param("token")

	response, err := accountService.ActivateUser(c, &account.ActivateUserRequest{
		Token: token,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t AuthController) ForgotPassword(c *gin.Context) {
	var forgotPasswordRequest requests.ForgotPassword

	if err := c.BindJSON(&forgotPasswordRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}
	response, err := accountService.ForgotPassword(c, &account.ForgotPasswordRequest{
		Email: forgotPasswordRequest.Email,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t AuthController) ResetPassword(c *gin.Context) {
	token := c.Param("token")
	var resetPasswordRequest requests.ResetPassword

	if err := c.BindJSON(&resetPasswordRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}
	response, err := accountService.ResetPassword(c, &account.ResetPasswordRequest{
		Token:    token,
		Password: resetPasswordRequest.Password,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t AuthController) ChangePassword(c *gin.Context) {
	userID := c.MustGet(utils.JWT_USER_ID).(string)
	var changePasswordRequest requests.ChangePassword

	if err := c.BindJSON(&changePasswordRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := accountService.ChangePassword(c, &account.ChangePasswordRequest{
		UserID:      userID,
		OldPassword: changePasswordRequest.OldPassword,
		NewPassword: changePasswordRequest.NewPassword,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t AuthController) Self(c *gin.Context) {
	userID := c.MustGet(utils.JWT_USER_ID).(string)
	response, err := accountService.FindUserById(c, &account.FindUserByIdRequest{
		Id: userID,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}
