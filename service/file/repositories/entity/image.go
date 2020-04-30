package entity

type Image struct {
	BaseModel
	Name       string `gorm:"type:varchar(255)"`
	Path       string `gorm:"type:varchar(255)"`
	Slug       string `gorm:"type:varchar(255)"`
	Thumbnail  string `gorm:"type:varchar(255)"`
	Type       string `gorm:"type:varchar(255)"`
	Title      string `gorm:"type:varchar(255)"`
	Alt        string `gorm:"type:varchar(255)"`
	Decription string `gorm:"type:varchar(255)"`
	Info       string `gorm:"type:varchar(255)"`
}
