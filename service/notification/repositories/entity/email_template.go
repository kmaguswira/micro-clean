package entity

type EmailTemplate struct {
	BaseModel
	Title     string `gorm:"type:varchar(255)"`
	Subject   string `gorm:"type:varchar(255)"`
	HTML      string `gorm:"type:text"`
	PlainText string `gorm:"type:text"`
	Language  string `gorm:"type:varchar(255)"`
	FromName  string `gorm:"type:varchar(255)"`
	FromEmail string `gorm:"type:varchar(255)"`
}
