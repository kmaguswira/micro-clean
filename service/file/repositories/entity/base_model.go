package entity

import (
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `gorm:"type:char(20);primary_key"`
	CreatedAt time.Time      `gorm:"column:createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt"`
}

func (bm *BaseModel) BeforeCreate(*gorm.DB) (err error) {
	bm.ID = xid.New().String()
	return
}
