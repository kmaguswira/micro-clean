package entity

type ACL struct {
	BaseModel
	Handler   string `gorm:"type:varchar(255)"`
	IsPublic  bool   `gorm:"type:boolean"`
	Title     string `gorm:"type:varchar(255)"`
	Permitted string `gorm:"type:varchar(255)"`
}
