package repositories

import (
	"fmt"

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

	aclDomain := populateACLDomain(acl)

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
		return populateACLDomain(e)
	}).ToSlice(&result)

	return &result, nil
}

func (t *readRepository) FindACLByHandler(input string) (*domain.ACL, error) {
	var acl entity.ACL

	if err := t.db.Where("handler = ?", input).First(&acl).Error; err != nil {
		err := fmt.Errorf("ACL with handler %q not found", input)
		return nil, err
	}

	aclDomain := populateACLDomain(acl)

	return &aclDomain, nil
}

func (t *readRepository) FindRoleByID(input string) (*domain.Role, error) {
	var role entity.Role

	if err := t.db.Where("id = ?", input).First(&role).Error; err != nil {
		err := fmt.Errorf("Role with id %q not found", input)
		return nil, err
	}

	roleDomain := populateRoleDomain(role)

	return &roleDomain, nil
}

func (t *readRepository) FindRoleByTitle(input string) (*domain.Role, error) {
	var role entity.Role
	if err := t.db.Where("title = ?", input).First(&role).Error; err != nil {
		err := fmt.Errorf("Role with title %q not found", input)
		return nil, err
	}

	roleDomain := populateRoleDomain(role)
	return &roleDomain, nil
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
		return populateRoleDomain(e)
	}).ToSlice(&result)

	return &result, nil
}

func (t *readRepository) FindUserByID(input string) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("id = ?", input).First(&user).Error; err != nil {
		err := fmt.Errorf("User with id %q not found", input)
		return nil, err
	}

	userDomain := populateUserDomain(user)
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
		user := populateUserDomain(e)
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

	userDomain := populateUserDomain(user)

	return &userDomain, nil
}

func (t *readRepository) FindUserByActivationToken(input string) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("activation_token = ?", input).First(&user).Error; err != nil {
		err := fmt.Errorf("User with activation token %q not found", input)
		return nil, err
	}

	userDomain := populateUserDomain(user)

	return &userDomain, nil
}

func (t *readRepository) FindUserByResetPasswordToken(input string) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("forgot_password_token = ?", input).First(&user).Error; err != nil {
		err := fmt.Errorf("User with forgot password token %q not found", input)
		return nil, err
	}

	userDomain := populateUserDomain(user)

	return &userDomain, nil
}
