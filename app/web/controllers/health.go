package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/utils"
)

type HealthController struct {
	utils.Response
}

func (t HealthController) Check(c *gin.Context) {
	t.OKSingleData(c, "OK")
	return
}
