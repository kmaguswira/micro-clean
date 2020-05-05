package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/domain"
	"github.com/kmaguswira/micro-clean/service/account/repositories"
	"github.com/kmaguswira/micro-clean/service/account/repositories/seeder/config"
)

func main() {
	env := flag.String("e", "local", "")

	flag.Usage = func() {
		fmt.Println("FLAG	DESCRIPTION		DEFAULT			OPTIONS")
		fmt.Println("-e		Environment		local			development/local/local_test/staging/production/test/migrate")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*env)

	seeding()
}

func seeding() {
	readWriteRepository := repositories.NewReadWriteRepository(nil)

	sr, _ := readWriteRepository.CreateRole(&SuperAdminRole)
	ar, _ := readWriteRepository.CreateRole(&AdminRole)
	ur, _ := readWriteRepository.CreateRole(&UserRole)

	readWriteRepository.CreateUser(SuperAdminUser(sr.ID))
	readWriteRepository.CreateUser(AdminUser(ar.ID))
	readWriteRepository.CreateUser(User(ur.ID))

	permitted := []domain.Role{*sr, *ar, *ur}
	acl := ACLs(permitted)

	for i := 0; i < len(acl); i++ {
		readWriteRepository.CreateACL(&acl[i])
	}
}

var SuperAdminRole = domain.Role{Title: global.SUPERADMIN_ROLE}
var AdminRole = domain.Role{Title: global.ADMIN_ROLE}
var UserRole = domain.Role{Title: global.USER_ROLE}

var SuperAdminUser = func(role string) *domain.User {
	d := domain.User{
		Email:    "superadmin@mail.com",
		Username: "superadmin",
		Name:     "superadmin",
		RoleID:   role,
		Status:   global.VERIFIED_USER_STATUS,
	}

	d.SetPassword("Password1$")
	return &d
}

var AdminUser = func(role string) *domain.User {
	d := domain.User{
		Email:    "admin@mail.com",
		Username: "admin",
		Name:     "admin",
		RoleID:   role,
		Status:   global.VERIFIED_USER_STATUS,
	}

	d.SetPassword("Password1$")
	return &d
}

var User = func(role string) *domain.User {
	d := domain.User{
		Email:    "user@mail.com",
		Username: "user",
		Name:     "user",
		RoleID:   role,
		Status:   global.VERIFIED_USER_STATUS,
	}

	d.SetPassword("Password1$")
	return &d
}

var ACLs = func(permitted []domain.Role) []domain.ACL {
	return []domain.ACL{
		domain.ACL{
			Handler:   "github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.HealthController.Check-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.AuthController.SignUp-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.AuthController.SignIn-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.AuthController.ActivateUser-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.AuthController.ForgotPassword-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.AuthController.ResetPassword-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.AuthController.ChangePassword-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.AuthController.Self-fm",
			Title:     "",
			IsPublic:  false,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.UserController.Create-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.UserController.FindById-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.UserController.Delete-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.UserController.Update-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.UserController.FindAll-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.RoleController.Create-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.RoleController.FindById-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.RoleController.Delete-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.RoleController.Update-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.RoleController.FindAll-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.ACLController.Create-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.ACLController.FindById-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.ACLController.Delete-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.ACLController.Update-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.ACLController.FindAll-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.EmailTemplateController.Create-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.EmailTemplateController.FindById-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.EmailTemplateController.Delete-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.EmailTemplateController.Update-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.EmailTemplateController.FindAll-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
		domain.ACL{
			Handler:   "github.com/kmaguswira/micro-clean/app/web/controllers.EmailTemplateController.SendEmail-fm",
			Title:     "",
			IsPublic:  true,
			Permitted: permitted,
		},
	}

}
