package entity

import (
	"time"

	"github.com/rs/xid"
)

type BaseModel struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (bm *BaseModel) BeforeCreate() (err error) {
	bm.ID = xid.New().String()
	return
}
