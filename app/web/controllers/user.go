package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
)

type UserController struct {
	utils.Response
}

func (t UserController) Create(c *gin.Context) {
	var createUserRequest requests.CreateUser

	if err := c.BindJSON(&createUserRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := accountService.CreateUser(c, &account.CreateUserRequest{
		New: &account.User{
			Name:     createUserRequest.Name,
			Username: createUserRequest.Username,
			Email:    createUserRequest.Email,
			RoleId:   createUserRequest.RoleID,
			Status:   createUserRequest.Status,
			Password: createUserRequest.Password,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t UserController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := accountService.FindUserById(c, &account.FindUserByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := accountService.DeleteUser(c, &account.DeleteUserRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}
