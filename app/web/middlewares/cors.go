package middlewares

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/config"
)

type corsMiddleware struct{}

func NewCORSMiddleware() corsMiddleware {
	return corsMiddleware{}
}

func (t *corsMiddleware) Handler() gin.HandlerFunc {
	config := config.GetConfig()

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", strings.Join(config.Security.Cors.AllowOrigin, ","))
		c.Writer.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.Security.Cors.MaxAge))
		c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(config.Security.Cors.AllowMethod, ","))
		c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(config.Security.Cors.AllowHeaders, ","))
		c.Writer.Header().Set("Access-Control-Expose-Headers", strings.Join(config.Security.Cors.ExposeHeaders, ","))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.Security.Cors.AllowCredentials))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}

}
