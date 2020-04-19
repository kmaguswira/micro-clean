package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name            string `gorm:"type:varchar(255)"`
	Email           string `gorm:"type:varchar(255)"`
	Username        string `gorm:"type:varchar(255)"`
	Password        string `gorm:"type:varchar(255)"`
	ActivationToken string `gorm:"type:varchar(255)"`
	RefreshToken    string `gorm:"type:varchar(255)"`
	RoleID          string `gorm:"type:char(20) REFERENCES roles(id)"`
	Role            Role   `gorm:"foreignkey:ID;association_foreignkey:RoleID"`
	Status          string `gorm:"type:varchar(255)"`
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
