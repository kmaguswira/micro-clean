package repositories

import (
	"github.com/kmaguswira/micro-clean/service/account/repositories/entity"
)

var Registered = []interface{}{
	entity.ACL{},
	entity.Role{},
	entity.User{},
}
