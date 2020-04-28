package repositories

import (
	"github.com/kmaguswira/micro-clean/service/account/application/global"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IReadRepository interface {
	FindACLByID(input string) (*domain.ACL, error)
	FindAllACL(input global.FindAllInput) (*[]domain.ACL, error)
	FindRoleByID(input string) (*domain.Role, error)
	FindAllRole(input global.FindAllInput) (*[]domain.Role, error)
	FindRoleByTitle(input string) (*domain.Role, error)
	FindUserByID(input string) (*domain.User, error)
	FindAllUser(input global.FindAllInput) (*[]domain.User, error)
	FindUserByEmailOrUsername(input string) (*domain.User, error)
	FindUserByActivationToken(input string) (*domain.User, error)
	FindUserByResetPasswordToken(input string) (*domain.User, error)
}
