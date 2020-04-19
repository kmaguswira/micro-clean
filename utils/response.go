package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct{}

func (r Response) OKSingleData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   data,
	})
}

func (r Response) OKMultipleData(c *gin.Context, data interface{}, total int, offset int) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   data,
		"total":  total,
		"offset": offset,
	})
}

func (r Response) BadRequest(c *gin.Context, info string) {
	message := "Your request is not valid, please check again"
	if info != "" {
		message = info
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  400,
		"message": message,
	})
}

func (r Response) NotAuthorize(c *gin.Context, info string) {
	message := "You're not authorized yet, please logged in"
	if info != "" {
		message = info
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  401,
		"message": message,
	})
}

func (r Response) Forbidden(c *gin.Context, info string) {
	message := "You don't have permission to access this data"
	if info != "" {
		message = info
	}

	c.JSON(http.StatusForbidden, gin.H{
		"status":  403,
		"message": message,
	})
}

func (r Response) NotFound(c *gin.Context, info string) {
	message := "Data that you're looking for is not found"
	if info != "" {
		message = info
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": message,
	})
}

func (r Response) InternalError(c *gin.Context, info string) {
	message := "Something error with our system"
	if info != "" {
		message = info
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
