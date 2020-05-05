package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/utils"
)

type healthController struct {
	utils.Response
}

func NewHealthController() healthController {
	return healthController{}
}

func (t *healthController) Check(c *gin.Context) {
	test := utils.QueryBuilder(c)
	t.OKSingleData(c, test)
	return
}
