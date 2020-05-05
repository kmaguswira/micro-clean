package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro/client"
)

type userController struct {
	utils.Response
	accountService account.AccountService
}

func NewUserController(client client.Client) userController {
	return userController{
		accountService: account.NewAccountService("kmaguswira.srv.account", client),
	}
}

func (t *userController) Create(c *gin.Context) {
	var createUserRequest requests.CreateUser

	if err := c.BindJSON(&createUserRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.CreateUser(c, &account.CreateUserRequest{
		New: &account.User{
			Name:     createUserRequest.Name,
			Username: createUserRequest.Username,
			Email:    createUserRequest.Email,
			RoleID:   createUserRequest.RoleID,
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

func (t *userController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := t.accountService.FindUserById(c, &account.FindUserByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *userController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := t.accountService.DeleteUser(c, &account.DeleteUserRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *userController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateUserRequest requests.UpdateUser

	if err := c.BindJSON(&updateUserRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.UpdateUser(c, &account.UpdateUserRequest{
		Update: &account.User{
			ID:       id,
			Name:     updateUserRequest.Name,
			Username: updateUserRequest.Username,
			Email:    updateUserRequest.Email,
			RoleID:   updateUserRequest.RoleID,
			Status:   updateUserRequest.Status,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *userController) FindAll(c *gin.Context) {
	q := utils.QueryBuilder(c)

	response, err := t.accountService.FindAllUser(c, &account.FindAllUserRequest{
		Query: &account.BaseQuery{
			QueryKey:   q.QueryKey,
			QueryValue: q.QueryValue,
			Limit:      q.Limit,
			Offset:     q.Offset,
			Sort:       q.Sort,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}
