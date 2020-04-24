package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/requests"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
)

type ACLController struct {
	utils.Response
}

func (t ACLController) Create(c *gin.Context) {
	var createACLRequest requests.CreateACL

	if err := c.BindJSON(&createACLRequest); err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	fmt.Println(createACLRequest.IsPublic)

	response, err := accountService.CreateACL(c, &account.CreateACLRequest{
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

func (t ACLController) FindById(c *gin.Context) {
	id := c.Param("id")

	response, err := accountService.FindACLById(c, &account.FindACLByIdRequest{
		Id: id,
	})

	if err != nil {
		t.BadRequest(c, err.Error())
		return
	}

	t.OKSingleData(c, response)

	return
}
