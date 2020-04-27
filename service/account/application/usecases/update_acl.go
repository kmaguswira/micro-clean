package usecases

import (
	"errors"
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type IUpdateACL interface {
	Execute(input UpdateACLInput) (domain.ACL, error)
}

type UpdateACLInput struct {
	ID        string
	Title     string
	Handler   string
	IsPublic  bool
	Permitted string
}

type updateACLUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewUpdateACLUseCase(ReadWriteRepository repositories.IReadWriteRepository) IUpdateACL {
	return &updateACLUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *updateACLUseCase) Execute(input UpdateACLInput) (domain.ACL, error) {
	var role []domain.Role
	roleIds := strings.Split(input.Permitted, ",")

	From(roleIds).Select(func(c interface{}) interface{} {
		return domain.Role{
			ID: c.(string),
		}
	}).ToSlice(&role)

	acl := domain.ACL{
		ID:        input.ID,
		Title:     input.Title,
		Handler:   input.Handler,
		IsPublic:  input.IsPublic,
		Permitted: role,
	}

	result, err := t.readWriteRepository.UpdateACL(&acl)

	if err == nil {
		return *result, nil
	}

	return domain.ACL{}, errors.New("Bad Request")
}
