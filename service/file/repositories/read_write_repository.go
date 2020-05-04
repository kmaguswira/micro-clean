package repositories

import (
	"github.com/jinzhu/gorm"
	iface "github.com/kmaguswira/micro-clean/service/file/application/repositories"
	"github.com/kmaguswira/micro-clean/service/file/config"
)

type readWriteRepository struct {
	db *gorm.DB
}

func NewReadWriteRepository(DB *gorm.DB) iface.IReadWriteRepository {
	if DB != nil {
		return &readWriteRepository{
			db: DB,
		}
	}
	config := config.GetConfig()
	return &readWriteRepository{
		db: NewDB(config.Repositories.ReadWrite, Registered),
	}
}
