package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/config"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type authController struct {
	accountService     account.AccountService
	sendEmailPublisher micro.Publisher
	utils.Response
}

func NewAuthController(client client.Client) authController {
	return authController{
		accountService:     account.NewAccountService("kmaguswira.srv.account", client),
		sendEmailPublisher: micro.NewPublisher("kmaguswira.srv.notification.send-email", client),
	}
}

func (t *authController) SignIn(c *gin.Context) {
	var signInRequest requests.SignIn

	if err := c.BindJSON(&signInRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.SignIn(c, &account.SignInRequest{
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

func (t *authController) SignUp(c *gin.Context) {
	var signUpRequest requests.SignUp

	if err := c.BindJSON(&signUpRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.SignUp(c, &account.SignUpRequest{
		Email:    signUpRequest.Email,
		Username: signUpRequest.Username,
		Password: signUpRequest.Password,
		Name:     signUpRequest.Name,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	//TODO: send email activation
	ev := &notification.SendEmailRequest{
		ToName:        signUpRequest.Name,
		ToEmail:       signUpRequest.Email,
		TemplateTitle: "activation",
		Data:          []string{signUpRequest.Name},
	}
	if err := t.sendEmailPublisher.Publish(c, ev); err != nil {
		log.Printf("error publishing: %v", err)
	}

	t.OKSingleData(c, response)
	return
}

func (t *authController) ActivateUser(c *gin.Context) {
	token := c.Param("token")

	response, err := t.accountService.ActivateUser(c, &account.ActivateUserRequest{
		Token: token,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t *authController) ForgotPassword(c *gin.Context) {
	var forgotPasswordRequest requests.ForgotPassword

	if err := c.BindJSON(&forgotPasswordRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}
	response, err := t.accountService.ForgotPassword(c, &account.ForgotPasswordRequest{
		Email: forgotPasswordRequest.Email,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t *authController) ResetPassword(c *gin.Context) {
	token := c.Param("token")
	var resetPasswordRequest requests.ResetPassword

	if err := c.BindJSON(&resetPasswordRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}
	response, err := t.accountService.ResetPassword(c, &account.ResetPasswordRequest{
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

func (t *authController) ChangePassword(c *gin.Context) {
	userID := c.MustGet(utils.JWT_USER_ID).(string)
	var changePasswordRequest requests.ChangePassword

	if err := c.BindJSON(&changePasswordRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.ChangePassword(c, &account.ChangePasswordRequest{
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

func (t *authController) Self(c *gin.Context) {
	userID := c.MustGet(utils.JWT_USER_ID).(string)
	response, err := t.accountService.FindUserById(c, &account.FindUserByIdRequest{
		Id: userID,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}
