package controllers

import (
	"github.com/asim/go-micro/v3/client"
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
)

type roleController struct {
	utils.Response
	accountService account.AccountService
}

func NewRoleController(client client.Client) roleController {
	return roleController{
		accountService: account.NewAccountService("kmaguswira.srv.account", client),
	}
}

func (t *roleController) Create(c *gin.Context) {
	var createRoleRequest requests.CreateRole

	if err := c.BindJSON(&createRoleRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.CreateRole(c, &account.CreateRoleRequest{
		New: &account.Role{
			Title: createRoleRequest.Title,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t *roleController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := t.accountService.FindRoleById(c, &account.FindRoleByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *roleController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := t.accountService.DeleteRole(c, &account.DeleteRoleRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *roleController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateRoleRequest requests.UpdateRole

	if err := c.BindJSON(&updateRoleRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.UpdateRole(c, &account.UpdateRoleRequest{
		Update: &account.Role{
			ID:    id,
			Title: updateRoleRequest.Title,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *roleController) FindAll(c *gin.Context) {
	q := utils.QueryBuilder(c)

	response, err := t.accountService.FindAllRole(c, &account.FindAllRoleRequest{
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
