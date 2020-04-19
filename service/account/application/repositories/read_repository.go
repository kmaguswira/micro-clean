package repositories

import (
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IReadRepository interface {
	FindRoleByTitle(input string) (*domain.Role, error)
	FindUserByEmailOrUsername(input string) (*domain.User, error)
}
