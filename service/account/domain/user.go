package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Name     string
	Username string
	Email    string
	Role     Role
	RoleID   string
	Status   string
	password string
}

func (t *User) SetPassword(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	t.password = string(bytes)
}

func (t *User) GetPassword() string {
	return t.password
}

func (t *User) SetHashedPassword(password string) {
	t.password = password
}

func (t *User) CheckHashedPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(t.password), []byte(password))

	return err == nil
}
