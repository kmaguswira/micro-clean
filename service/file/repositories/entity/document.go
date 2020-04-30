package entity

type Document struct {
	BaseModel
	Name       string `gorm:"type:varchar(255)"`
	Path       string `gorm:"type:varchar(255)"`
	Slug       string `gorm:"type:varchar(255)"`
	Type       string `gorm:"type:varchar(255)"`
	Decription string `gorm:"type:varchar(255)"`
	Info       string `gorm:"type:varchar(255)"`
}
