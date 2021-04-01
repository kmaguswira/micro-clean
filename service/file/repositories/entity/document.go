package entity

type Document struct {
	BaseModel
	Name        string `gorm:"type:varchar(255)"`
	Path        string `gorm:"type:text"`
	Slug        string `gorm:"type:text"`
	Type        string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Info        string `gorm:"type:varchar(255)"`
	Cdn         bool   `gorm:"type:boolean"`
}
