package controllers

import (
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

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}
	cfg := config.GetConfig()
	token := utils.GenerateToken(response.Result.Id, response.Result.Role, cfg.Server.Secret)

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
