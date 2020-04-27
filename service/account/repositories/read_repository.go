package repositories

import (
	"fmt"
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	"github.com/kmaguswira/micro-clean/service/account/application/global"
	iface "github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/config"
	"github.com/kmaguswira/micro-clean/service/account/domain"
	"github.com/kmaguswira/micro-clean/service/account/repositories/entity"
)

type readRepository struct {
	db *gorm.DB
}

func NewReadRepository(DB *gorm.DB) iface.IReadRepository {
	if DB != nil {
		return &readRepository{
			db: DB,
		}
	}
	config := config.GetConfig()
	return &readRepository{
		db: NewDB(config.Repositories.Read, Registered),
	}
}

func (t *readRepository) FindACLByID(input string) (*domain.ACL, error) {
	var acl entity.ACL

	if err := t.db.Where("id = ?", input).First(&acl).Error; err != nil {
		err := fmt.Errorf("ACL with id %q not found", input)
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
		Handler:   acl.Handler,
		IsPublic:  acl.IsPublic,
		Title:     acl.Title,
		Permitted: role,
	}

	return &aclDomain, nil
}

func (t *readRepository) FindAllACL(input global.FindAllInput) (*[]domain.ACL, error) {
	var acls []entity.ACL
	var result []domain.ACL

	if err := t.db.Order(input.Sort).Limit(input.Limit).Offset(input.Offset).Where(input.QueryKey, input.QueryValue...).Find(&acls).Error; err != nil {
		err := fmt.Errorf("Query Error", input)
		return nil, err
	}

	From(acls).Select(func(c interface{}) interface{} {
		e := c.(entity.ACL)
		var roles []domain.Role
		roleIds := strings.Split(e.Permitted, ",")

		From(roleIds).Select(func(c interface{}) interface{} {
			return domain.Role{
				ID: c.(string),
			}
		}).ToSlice(&roles)

		acl := domain.ACL{
			ID:        e.ID,
			Title:     e.Title,
			IsPublic:  e.IsPublic,
			Permitted: roles,
			Handler:   e.Handler,
		}

		return acl
	}).ToSlice(&result)

	return &result, nil
}

func (t *readRepository) FindRoleByID(input string) (*domain.Role, error) {
	var role entity.Role

	if err := t.db.Where("id = ?", input).First(&role).Error; err != nil {
		err := fmt.Errorf("Role with id %q not found", input)
		return nil, err
	}

	roleDomain := domain.Role{
		ID:    role.ID,
		Title: role.Title,
	}

	return &roleDomain, nil
}

func (t *readRepository) FindRoleByTitle(input string) (*domain.Role, error) {
	var role domain.Role
	if err := t.db.Where("title = ?", input).First(&role).Error; err != nil {
		err := fmt.Errorf("Role with title %q not found", input)
		return nil, err
	}
	return &role, nil
}

func (t *readRepository) FindAllRole(input global.FindAllInput) (*[]domain.Role, error) {
	var roles []entity.Role
	var result []domain.Role

	if err := t.db.Order(input.Sort).Limit(input.Limit).Offset(input.Offset).Where(input.QueryKey, input.QueryValue...).Find(&roles).Error; err != nil {
		err := fmt.Errorf("Query Error", input)
		return nil, err
	}

	From(roles).Select(func(c interface{}) interface{} {
		e := c.(entity.Role)
		return domain.Role{
			ID:    e.ID,
			Title: e.Title,
		}
	}).ToSlice(&result)

	return &result, nil
}

func (t *readRepository) FindUserByID(input string) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("id = ?", input).First(&user).Error; err != nil {
		err := fmt.Errorf("User with id %q not found", input)
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
	return &userDomain, nil
}

func (t *readRepository) FindAllUser(input global.FindAllInput) (*[]domain.User, error) {
	var users []entity.User
	var result []domain.User

	if err := t.db.Order(input.Sort).Limit(input.Limit).Offset(input.Offset).Where(input.QueryKey, input.QueryValue...).Find(&users).Error; err != nil {
		err := fmt.Errorf("Query Error", input)
		return nil, err
	}

	From(users).Select(func(c interface{}) interface{} {
		e := c.(entity.User)
		user := domain.User{
			ID:       e.ID,
			Name:     e.Name,
			Username: e.Username,
			Email:    e.Email,
			Status:   e.Status,
			RoleID:   e.RoleID,
		}

		user.SetHashedPassword(e.Password)
		return user
	}).ToSlice(&result)

	return &result, nil
}

func (t *readRepository) FindUserByEmailOrUsername(input string) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("email = ?", input).Or("username = ?", input).First(&user).Error; err != nil {
		err := fmt.Errorf("User with email or username %q not found", input)
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

	return &userDomain, nil
}
