package repo

import (
	"github.com/jinzhu/gorm"
)

type SettingRepo interface {
	GetNextId(uint32) (uint32, error)
}

type SettingRepoImpl struct {
	DB *gorm.DB
}

// Version 1
func (self *SettingRepoImpl) GetNextId_V1(increStep uint32) (uint32, error) {
	var valueInt uint32
	self.DB.Raw("SELECT value_int FROM settings WHERE `key`=? LIMIT 1 FOR UPDATE", "increment_id").
		Row().
		Scan(&valueInt)

	valueInt += increStep

	self.DB.Exec("UPDATE settings SET value_int = ? WHERE `key`= ? LIMIT 1", valueInt, "increment_id")
	return valueInt, nil
}

// Version 2
// Voi lock row
func (self *SettingRepoImpl) GetNextId(increStep uint32) (uint32, error) {
	tx := self.DB.Begin()
	var valueInt uint32
	tx.Raw("SELECT value_int FROM settings WHERE `key`=? LIMIT 1 FOR UPDATE", "increment_id").
		Row().
		Scan(&valueInt)

	// Nen tang mot luc n don vi
	valueInt += increStep

	// Chu y ban co the Sleep de test row bi lock nhu the nao
	// time.Sleep(time.Second * 10)

	tx.Exec("UPDATE settings SET value_int = ? WHERE `key`= ? LIMIT 1", valueInt, "increment_id")
	tx.Commit()
	return valueInt, nil
}
