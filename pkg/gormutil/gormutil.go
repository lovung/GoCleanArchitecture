package gormutil

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once      sync.Once
	singleton *gorm.DB
)

// OpenDBConnection opens a DB connection.
func OpenDBConnection(connStr string, config gorm.Config) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(connStr), &config)
		if err == nil {
			singleton = db
		}
	})
	if err != nil {
		return nil, err
	}
	return singleton, nil
}

// GetDB gets the instance of singleton
func GetDB() *gorm.DB {
	return singleton
}
