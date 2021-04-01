package repositories

import (
	"fmt"
	"log"

	// "gorm.io/driver/sqlite"
	// "gorm.io/driver/postgres"
	"github.com/kmaguswira/micro-clean/service/account/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db  *gorm.DB
	err error
}

func NewDB(connection config.DB) *gorm.DB {
	db := new(DB)
	db.Connect(connection)

	return db.db
}

func (t *DB) Connect(connection config.DB) {
	dbinfo := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		connection.User,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.Name,
	)

	if connection.Driver == "postgre" {
		dbinfo = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			connection.User,
			connection.Password,
			connection.Host,
			connection.Port,
			connection.Name,
		)
	}

	t.db, t.err = gorm.Open(mysql.Open(dbinfo), &gorm.Config{})
	if t.err != nil {
		log.Println("Failed to connect to database")
		panic(t.err)
	}

	// t.db.LogMode(true)
	log.Println("DB connected")
}
