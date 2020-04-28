package handler

import (
	"context"
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/kmaguswira/micro-clean/service/account/application/usecases"
	"github.com/kmaguswira/micro-clean/service/account/domain"
	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
	"github.com/kmaguswira/micro-clean/service/account/repositories"
	"github.com/kmaguswira/micro-clean/service/account/utils"
	"github.com/micro/go-micro/util/log"
)

type Account struct {
	signUpUseCase          usecases.ISignUp
	signInUseCase          usecases.ISignIn
	createRoleUseCase      usecases.ICreateRole
	findRoleByIDUseCase    usecases.IFindRoleByID
	deleteRoleUseCase      usecases.IDeleteRole
	updateRoleUseCase      usecases.IUpdateRole
	findAllRoleUseCase     usecases.IFindAllRole
	findRoleByTitleUseCase usecases.IFindRoleByTitle
	createACLUseCase       usecases.ICreateACL
	findACLByIDUseCase     usecases.IFindACLByID
	deleteACLUseCase       usecases.IDeleteACL
	updateACLUseCase       usecases.IUpdateACL
	findAllACLUseCase      usecases.IFindAllACL
	createUserUseCase      usecases.ICreateUser
	findUserByIDUseCase    usecases.IFindUserByID
	deleteUserUseCase      usecases.IDeleteUser
	updateUserUseCase      usecases.IUpdateUser
	findAllUserUseCase     usecases.IFindAllUser
	activateUserUseCase    usecases.IActivateUser
	forgotPasswordUseCase  usecases.IForgotPassword
	resetPasswordUseCase   usecases.IResetPassword
	changePasswordUseCase  usecases.IChangePassword
	response               utils.Response
}

func NewAccount() *Account {
	readWriteRepository := repositories.NewReadWriteRepository(nil)
	readRepository := repositories.NewReadRepository(nil)

	return &Account{
		signUpUseCase:          usecases.NewSignUpUseCase(readWriteRepository),
		signInUseCase:          usecases.NewSignInUseCase(readRepository),
		createRoleUseCase:      usecases.NewCreateRoleUseCase(readWriteRepository),
		findRoleByIDUseCase:    usecases.NewFindRoleByIDUseCase(readRepository),
		deleteRoleUseCase:      usecases.NewDeleteRoleUseCase(readWriteRepository),
		updateRoleUseCase:      usecases.NewUpdateRoleUseCase(readWriteRepository),
		findAllRoleUseCase:     usecases.NewFindAllRoleUseCase(readRepository),
		findRoleByTitleUseCase: usecases.NewfindRoleByTitleUseCase(readRepository),
		createACLUseCase:       usecases.NewCreateACLUseCase(readWriteRepository),
		findACLByIDUseCase:     usecases.NewFindACLByIDUseCase(readRepository),
		deleteACLUseCase:       usecases.NewDeleteACLUseCase(readWriteRepository),
		updateACLUseCase:       usecases.NewUpdateACLUseCase(readWriteRepository),
		findAllACLUseCase:      usecases.NewFindAllACLUseCase(readRepository),
		createUserUseCase:      usecases.NewCreateUserUseCase(readWriteRepository),
		findUserByIDUseCase:    usecases.NewFindUserByIDUseCase(readRepository),
		deleteUserUseCase:      usecases.NewDeleteUserUseCase(readWriteRepository),
		updateUserUseCase:      usecases.NewUpdateUserUseCase(readWriteRepository),
		findAllUserUseCase:     usecases.NewFindAllUserUseCase(readRepository),
		activateUserUseCase:    usecases.NewActivateUserUseCase(readRepository, readWriteRepository),
		forgotPasswordUseCase:  usecases.NewForgotPasswordUseCase(readRepository, readWriteRepository),
		resetPasswordUseCase:   usecases.NewResetPasswordUseCase(readRepository, readWriteRepository),
		changePasswordUseCase:  usecases.NewChangePasswordUseCase(readRepository, readWriteRepository),
		response:               utils.Response{},
	}
}

// Call is a single request handler called via client.Call or the generated client code
func (t *Account) Call(ctx context.Context, req *account.Request, rsp *account.Response) error {
	log.Log("Received Account.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (t *Account) Stream(ctx context.Context, req *account.StreamingRequest, stream account.Account_StreamStream) error {
	log.Logf("Received Account.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&account.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (t *Account) PingPong(ctx context.Context, stream account.Account_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&account.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func populateUserResponse(user domain.User) account.User {
	userResponse := account.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Status:   user.Status,
		RoleID:   user.RoleID,
	}

	return userResponse
}

func populateRoleResponse(role domain.Role) account.Role {
	roleResponse := account.Role{
		ID:    role.ID,
		Title: role.Title,
	}

	return roleResponse
}

func populateACLResponse(acl domain.ACL) account.ACL {
	var roleID []string
	From(acl.Permitted).Select(func(c interface{}) interface{} {
		return c.(domain.Role).ID
	}).ToSlice(&roleID)

	aclResponse := account.ACL{
		ID:        acl.ID,
		Title:     acl.Title,
		Handler:   acl.Handler,
		IsPublic:  acl.IsPublic,
		Permitted: strings.Join(roleID, ","),
	}
	return aclResponse
}
