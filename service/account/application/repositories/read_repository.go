package repositories

import (
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IReadRepository interface {
	FindACLByID(input string) (*domain.ACL, error)
	FindRoleByID(input string) (*domain.Role, error)
	FindRoleByTitle(input string) (*domain.Role, error)
	FindUserByID(input string) (*domain.User, error)
	FindUserByEmailOrUsername(input string) (*domain.User, error)
}
