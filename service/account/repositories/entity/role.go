package entity

type Role struct {
	BaseModel
	Title string `gorm:"type:varchar(255)"`
	Users []User
}
