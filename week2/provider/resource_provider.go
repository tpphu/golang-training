package provider

import (
	"github.com/tpphu/golang-training/week2/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RP struct {
	DB *gorm.DB
}

func MustBuildRP(conf *config.Config) *RP {
	dsn := conf.MySQLDNS
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	return &RP{
		DB: db,
	}
}
