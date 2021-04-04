package entity

type ACL struct {
	BaseModel
	Handler   string `gorm:"type:text"`
	IsPublic  bool   `gorm:"type:boolean;column:isPublic"`
	Title     string `gorm:"type:varchar(255)"`
	Permitted string `gorm:"type:text"`
}
