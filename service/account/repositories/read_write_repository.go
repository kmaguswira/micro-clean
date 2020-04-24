package repositories

import (
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	iface "github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/config"
	"github.com/kmaguswira/micro-clean/service/account/domain"
	"github.com/kmaguswira/micro-clean/service/account/repositories/entity"
)

type readWriteRepository struct {
	db *gorm.DB
}

func NewReadWriteRepository(DB *gorm.DB) iface.IReadWriteRepository {
	if DB != nil {
		return &readWriteRepository{
			db: DB,
		}
	}
	config := config.GetConfig()
	return &readWriteRepository{
		db: NewDB(config.Repositories.Read, Registered),
	}
}

func (t *readWriteRepository) CreateRole(role *domain.Role) (*domain.Role, error) {
	newRole := entity.Role{
		Title: role.Title,
	}
	t.db.Create(&newRole)
	return role, nil
}

func (t *readWriteRepository) CreateACL(acl *domain.ACL) (*domain.ACL, error) {
	var roleID []string
	From(acl.Permitted).Select(func(c interface{}) interface{} {
		return c.(domain.Role).ID
	}).ToSlice(&roleID)

	newACL := entity.ACL{
		Title:     acl.Title,
		Handler:   acl.Handler,
		IsPublic:  acl.IsPublic,
		Permitted: strings.Join(roleID, ","),
	}

	t.db.Create(&newACL)
	return acl, nil
}

func (t *readWriteRepository) CreateUser(user *domain.User) (*domain.User, error) {
	newUser := entity.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
		Password: user.GetPassword(),
		RoleID:   user.RoleID,
	}
	t.db.Create(&newUser)
	user.ID = newUser.ID
	return user, nil
}
