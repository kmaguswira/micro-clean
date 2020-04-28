package repositories

import (
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/kmaguswira/micro-clean/service/account/domain"
	"github.com/kmaguswira/micro-clean/service/account/repositories/entity"
)

func populateUserDomain(user entity.User) domain.User {
	userDomain := domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		RoleID:   user.RoleID,
		Status:   user.Status,
	}

	userDomain.SetHashedPassword(user.Password)
	userDomain.SetActivationToken(user.ActivationToken)

	return userDomain
}

func populateRoleDomain(role entity.Role) domain.Role {
	roleDomain := domain.Role{
		ID:    role.ID,
		Title: role.Title,
	}

	return roleDomain
}

func populateACLDomain(acl entity.ACL) domain.ACL {
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

	return aclDomain
}
