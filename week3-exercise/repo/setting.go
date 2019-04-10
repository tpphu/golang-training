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
	// Cho nay cung dung Lock Row nhung ma vi khong co transaction
	// Nen no khong co work
	self.DB.Raw("SELECT value_int FROM settings WHERE `key`=? LIMIT 1 FOR UPDATE", "increment_id").
		Row().
		Scan(&valueInt)

	valueInt += increStep

	self.DB.Exec("UPDATE settings SET value_int = ? WHERE `key`= ? LIMIT 1", valueInt, "increment_id")
	return valueInt, nil
}

// Hien tai la current: 0, next_value: 100
// Hien tai la current: 100, next value: 200
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
