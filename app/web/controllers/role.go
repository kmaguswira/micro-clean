package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
)

type RoleController struct {
	utils.Response
}

func (t RoleController) Create(c *gin.Context) {
	var createRoleRequest requests.CreateRole

	if err := c.BindJSON(&createRoleRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := accountService.CreateRole(c, &account.CreateRoleRequest{
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

func (t RoleController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := accountService.FindRoleById(c, &account.FindRoleByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}
