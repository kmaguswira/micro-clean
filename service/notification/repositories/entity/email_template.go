package entity

type EmailTemplate struct {
	BaseModel
	Title    string `gorm:"type:varchar(255)"`
	Subject  string `gorm:"type:varchar(255)"`
	Body     string `gorm:"type:varchar(255)"`
	Language string `gorm:"type:varchar(255)"`
}
