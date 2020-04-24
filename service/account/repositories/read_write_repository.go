package repositories

import (
	"fmt"
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

func (t *readWriteRepository) DeleteRole(ID string) (*domain.Role, error) {
	var role entity.Role

	if err := t.db.Where("id = ?", ID).First(&role).Error; err != nil {
		err := fmt.Errorf("Role with ID %q not found", ID)
		return nil, err
	}

	roleDomain := domain.Role{
		ID:    role.ID,
		Title: role.Title,
	}

	t.db.Delete(&role)

	return &roleDomain, nil
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

func (t *readWriteRepository) DeleteACL(ID string) (*domain.ACL, error) {
	var acl entity.ACL

	if err := t.db.Where("id = ?", ID).First(&acl).Error; err != nil {
		err := fmt.Errorf("ACL with ID %q not found", ID)
		return nil, err
	}

	var role []domain.Role
	roleIds := strings.Split(acl.Permitted, ",")

	From(roleIds).Select(func(c interface{}) interface{} {
		return domain.Role{
			ID: c.(string),
		}
	}).ToSlice(&role)

	aclDomain := domain.ACL{
		ID:        acl.ID,
		Title:     acl.Title,
		Handler:   acl.Handler,
		IsPublic:  acl.IsPublic,
		Permitted: role,
	}

	t.db.Delete(&acl)

	return &aclDomain, nil
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

func (t *readWriteRepository) DeleteUser(ID string) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("id = ?", ID).First(&user).Error; err != nil {
		err := fmt.Errorf("User with ID %q not found", ID)
		return nil, err
	}

	userDomain := domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		RoleID:   user.RoleID,
		Status:   user.Status,
	}

	userDomain.SetHashedPassword(user.Password)

	t.db.Delete(&user)

	return &userDomain, nil
}
