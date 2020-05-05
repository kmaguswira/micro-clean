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

type authorizationMiddleware struct {
	accountService account.AccountService
	utils.Response
}

func NewAuthorizationMiddleware(client client.Client) authorizationMiddleware {
	return authorizationMiddleware{
		accountService: account.NewAccountService("kmaguswira.srv.account", client),
	}
}

func (t *authorizationMiddleware) Handler() gin.HandlerFunc {
	cfg := config.GetConfig()

	return func(c *gin.Context) {
		response, _ := t.accountService.FindACLByHandler(c, &account.FindACLByHandlerRequest{
			Handler: c.HandlerName(),
		})

		if response.Result == nil {
			fmt.Println(c.HandlerName())
			t.InternalError(c, "Can't Connect to Account Service")
			return
		}

		if !response.Result.IsPublic {
			authorizationHeader := c.Request.Header.Get("Authorization")
			var userID string
			var roleID string

			if authorizationHeader != "" {
				token := strings.Split(authorizationHeader, " ")

				if token[0] == "Bearer" {
					isAuthorized, user, role, err := utils.IsValidToken(token[1], cfg.Server.Secret)

					if err != nil {
						log.Println(err)
						t.NotAuthorize(c, err.Error())
						return
					}

					if isAuthorized {
						userID = user
						roleID = role
						c.Set(utils.JWT_USER_ID, userID)
						c.Set(utils.JWT_ROLE_ID, roleID)
					} else {
						t.NotAuthorize(c, err.Error())
						return
					}
				}
			}

			if roleID == "" || !strings.Contains(response.Result.Permitted, roleID) {
				t.Forbidden(c, "Role Not Allowed")
				return
			}
		}

		c.Next()
	}
}
