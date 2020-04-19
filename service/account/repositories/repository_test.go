package repositories

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	iface "github.com/kmaguswira/micro-clean/service/account/application/repositories"
	"github.com/kmaguswira/micro-clean/service/account/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	readRepository      iface.IReadRepository
	readWriteRepository iface.IReadWriteRepository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.readRepository = NewReadRepository(s.DB)
	s.readWriteRepository = NewReadWriteRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestFindRoleByTitle() {
	title := "test"

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `roles` WHERE (title = ?) ORDER BY `roles`.`id` ASC LIMIT 1")).
		WithArgs(title).
		WillReturnRows(sqlmock.NewRows([]string{"title"}).AddRow(title))

	res, err := s.readRepository.FindRoleByTitle(title)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), &domain.Role{Title: "test"}, res)
}

func (s *Suite) TestFindUserByEmailOrUsername() {
	input := "test"

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?) OR (username = ?)) ORDER BY `users`.`id` ASC LIMIT 1")).
		WithArgs(input, input).
		WillReturnRows(sqlmock.NewRows([]string{"email", "username"}).AddRow(input, input))

	res, err := s.readRepository.FindUserByEmailOrUsername(input)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), &domain.User{Username: "test", Email: "test"}, res)
}

func (s *Suite) TestCreateRole() {
	role := domain.Role{
		Title: "title",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO `roles`").WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := s.readWriteRepository.CreateRole(&role)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), &domain.Role{Title: "title"}, res)

}
