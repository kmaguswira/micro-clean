package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/utils"
)

type HealthController struct {
	utils.Response
}

func (t HealthController) Check(c *gin.Context) {
	test := utils.QueryBuilder(c)
	t.OKSingleData(c, test)
	return
}
