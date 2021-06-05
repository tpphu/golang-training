package provider

import (
	"phudt/week4/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Provider struct {
	DB *gorm.DB
}

func NewProvider(cfg *config.Config) Provider {
	db, err := gorm.Open(mysql.Open(cfg.MySQL.ToString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return Provider{
		DB: db,
	}
}
