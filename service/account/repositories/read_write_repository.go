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

	role.ID = newRole.ID

	return role, nil
}

func (t *readWriteRepository) DeleteRole(ID string) (*domain.Role, error) {
	var role entity.Role

	if err := t.db.Where("id = ?", ID).First(&role).Error; err != nil {
		err := fmt.Errorf("Role with ID %q not found", ID)
		return nil, err
	}

	roleDomain := populateRoleDomain(role)
	t.db.Delete(&role)

	return &roleDomain, nil
}

func (t *readWriteRepository) UpdateRole(roleUpdated *domain.Role) (*domain.Role, error) {
	var role entity.Role

	if err := t.db.Where("id = ?", roleUpdated.ID).First(&role).Error; err != nil {
		err := fmt.Errorf("Role with ID %q not found", roleUpdated.ID)
		return nil, err
	}

	role.Title = roleUpdated.Title

	t.db.Save(&role)
	return roleUpdated, nil
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

	acl.ID = newACL.ID

	return acl, nil
}

func (t *readWriteRepository) DeleteACL(ID string) (*domain.ACL, error) {
	var acl entity.ACL

	if err := t.db.Where("id = ?", ID).First(&acl).Error; err != nil {
		err := fmt.Errorf("ACL with ID %q not found", ID)
		return nil, err
	}

	aclDomain := populateACLDomain(acl)

	t.db.Delete(&acl)

	return &aclDomain, nil
}

func (t *readWriteRepository) UpdateACL(aclUpdated *domain.ACL) (*domain.ACL, error) {
	var acl entity.ACL

	if err := t.db.Where("id = ?", aclUpdated.ID).First(&acl).Error; err != nil {
		err := fmt.Errorf("ACL with ID %q not found", aclUpdated.ID)
		return nil, err
	}

	var roleID []string
	From(aclUpdated.Permitted).Select(func(c interface{}) interface{} {
		return c.(domain.Role).ID
	}).ToSlice(&roleID)

	acl.Title = aclUpdated.Title
	acl.Handler = aclUpdated.Handler
	acl.IsPublic = aclUpdated.IsPublic
	acl.Permitted = strings.Join(roleID, ",")

	t.db.Save(&acl)
	return aclUpdated, nil
}

func (t *readWriteRepository) CreateUser(user *domain.User) (*domain.User, error) {
	newUser := entity.User{
		Name:            user.Name,
		Username:        user.Username,
		Email:           user.Email,
		Status:          user.Status,
		Password:        user.GetPassword(),
		RoleID:          user.RoleID,
		ActivationToken: user.GetActivationToken(),
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

	userDomain := populateUserDomain(user)
	t.db.Delete(&user)

	return &userDomain, nil
}

func (t *readWriteRepository) UpdateUser(userUpdated *domain.User) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("id = ?", userUpdated.ID).First(&user).Error; err != nil {
		err := fmt.Errorf("User with ID %q not found", userUpdated.ID)
		return nil, err
	}

	user.Name = userUpdated.Name
	user.Username = userUpdated.Username
	user.Email = userUpdated.Email
	user.RoleID = userUpdated.RoleID
	user.Status = userUpdated.Status

	t.db.Save(&user)
	return userUpdated, nil
}

func (t *readWriteRepository) ActivateUser(ID string) (*domain.User, error) {
	var user entity.User

	if err := t.db.Where("id = ?", ID).First(&user).Error; err != nil {
		err := fmt.Errorf("User with ID %q not found", ID)
		return nil, err
	}

	user.Status = "active"

	t.db.Save(&user)

	userDomain := populateUserDomain(user)
	return &userDomain, nil
}
