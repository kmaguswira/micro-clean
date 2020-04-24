package repositories

import (
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IReadWriteRepository interface {
	CreateRole(role *domain.Role) (*domain.Role, error)
	DeleteRole(ID string) (*domain.Role, error)
	CreateACL(acl *domain.ACL) (*domain.ACL, error)
	DeleteACL(ID string) (*domain.ACL, error)
	CreateUser(user *domain.User) (*domain.User, error)
	DeleteUser(ID string) (*domain.User, error)
}
