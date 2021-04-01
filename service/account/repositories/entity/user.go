package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name                string `gorm:"type:varchar(255)"`
	Email               string `gorm:"type:varchar(255)"`
	Username            string `gorm:"type:varchar(255)"`
	Password            string `gorm:"type:varchar(255)"`
	ActivationToken     string `gorm:"type:varchar(255);column:activationToken"`
	ForgotPasswordToken string `gorm:"type:varchar(255);column:forgotPasswordToken"`
	RefreshToken        string `gorm:"type:varchar(255);column:refreshToken"`
	RoleID              string `gorm:"type:char(20);column:roleId"`
	Role                Role
	Status              string `gorm:"type:varchar(255)"`
}

func (t *User) HashPassword() bool {
	bytes, err := bcrypt.GenerateFromPassword([]byte(t.Password), 14)
	t.Password = string(bytes)
	return err == nil
}

func (t *User) CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
