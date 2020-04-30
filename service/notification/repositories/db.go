package repositories

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/jinzhu/gorm/dialects/sqlite"
	// "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kmaguswira/micro-clean/service/notification/config"
	"github.com/kmaguswira/micro-clean/service/notification/utils"
)

type DB struct {
	db  *gorm.DB
	err error
}

func NewDB(connection config.DB, RegisteredTables []interface{}) *gorm.DB {
	db := new(DB)
	db.Connect(connection)
	db.Migrate(connection.Drop, RegisteredTables)

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

	t.db, t.err = gorm.Open(connection.Driver, dbinfo)

	if t.err != nil {
		log.Println("Failed to connect to database")
		panic(t.err)
	}

	// t.db.LogMode(true)
	log.Println("DB connected")
}

func (t *DB) Migrate(isDrop bool, tables []interface{}) {
	reverseTables := utils.ReverseArray(tables)

	for _, model := range reverseTables {
		if isDrop {
			t.db.DropTableIfExists(model)
		}
	}

	for _, model := range tables {
		if err := t.db.AutoMigrate(model).Error; err != nil {
			log.Println(err)
		} else {
			log.Println("Auto migrating", reflect.TypeOf(model).Name(), "...")
		}
	}
}
