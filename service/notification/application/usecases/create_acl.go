package usecases

import (
	"errors"
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
)

type ICreateACL interface {
	Execute(input CreateACLInput) (domain.ACL, error)
}

type CreateACLInput struct {
	Title     string
	Handler   string
	IsPublic  bool
	Permitted string
}

type createACLUseCase struct {
	readWriteRepository repositories.IReadWriteRepository
}

func NewCreateACLUseCase(ReadWriteRepository repositories.IReadWriteRepository) ICreateACL {
	return &createACLUseCase{
		readWriteRepository: ReadWriteRepository,
	}
}

func (t *createACLUseCase) Execute(input CreateACLInput) (domain.ACL, error) {
	var role []domain.Role
	roleIds := strings.Split(input.Permitted, ",")

	From(roleIds).Select(func(c interface{}) interface{} {
		return domain.Role{
			ID: c.(string),
		}
	}).ToSlice(&role)

	newACL := domain.ACL{
		Title:     input.Title,
		Handler:   input.Handler,
		IsPublic:  input.IsPublic,
		Permitted: role,
	}
	result, err := t.readWriteRepository.CreateACL(&newACL)

	if err == nil {
		return *result, nil
	}
	return *result, errors.New("Bad Request")
}
