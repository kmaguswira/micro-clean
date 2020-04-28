package middlewares

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/config"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/utils"
	"github.com/micro/go-micro/client"
)

var accountService = account.NewAccountService("kmaguswira.srv.account", client.DefaultClient)

func AuthorizationMiddleware() gin.HandlerFunc {
	cfg := config.GetConfig()

	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		var userID string
		var roleID string

		if authorizationHeader != "" {
			token := strings.Split(authorizationHeader, " ")

			if token[0] == "Bearer" {
				isAuthorized, user, role, err := utils.IsValidToken(token[1], cfg.Server.Secret)

				if isAuthorized {
					userID = user
					roleID = role
					c.Set(utils.JWT_USER_ID, userID)
					c.Set(utils.JWT_ROLE_ID, roleID)
					fmt.Println(user, role)
				}

				if err != nil {
					log.Println(err)
				}
			}
		}

		fmt.Println(c.HandlerName())

		// Check acl
		// response, err := accountService.SignIn(c, &account.SignInRequest{
		// 	RoleID: RoleID,
		// 	Handler: c.HandlerName(),
		// })

		// if err != nil {
		// 	t.BadRequest(c, err.Error())
		// 	return
		// }

		c.Next()
	}

}
