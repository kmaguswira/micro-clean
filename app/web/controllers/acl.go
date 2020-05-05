package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro/client"
)

type aclController struct {
	utils.Response
	accountService account.AccountService
}

func NewACLController(client client.Client) aclController {
	return aclController{
		accountService: account.NewAccountService("kmaguswira.srv.account", client),
	}
}

func (t *aclController) Create(c *gin.Context) {
	var createACLRequest requests.CreateACL

	if err := c.BindJSON(&createACLRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	fmt.Println(createACLRequest.IsPublic)

	response, err := t.accountService.CreateACL(c, &account.CreateACLRequest{
		New: &account.ACL{
			Handler:   createACLRequest.Handler,
			IsPublic:  createACLRequest.IsPublic,
			Title:     createACLRequest.Title,
			Permitted: createACLRequest.Permitted,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)
	return
}

func (t *aclController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := t.accountService.FindACLById(c, &account.FindACLByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *aclController) Delete(c *gin.Context) {
	id := c.Param("id")

	response, err := t.accountService.DeleteACL(c, &account.DeleteACLRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *aclController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateACLRequest requests.UpdateACL

	if err := c.BindJSON(&updateACLRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	response, err := t.accountService.UpdateACL(c, &account.UpdateACLRequest{
		Update: &account.ACL{
			ID:        id,
			Handler:   updateACLRequest.Handler,
			IsPublic:  updateACLRequest.IsPublic,
			Title:     updateACLRequest.Title,
			Permitted: updateACLRequest.Permitted,
		},
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}

func (t *aclController) FindAll(c *gin.Context) {
	q := utils.QueryBuilder(c)

	response, err := t.accountService.FindAllACL(c, &account.FindAllACLRequest{
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
