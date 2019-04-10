package handler

import (
	"sync"

	"../repo"
	"github.com/gin-gonic/gin"
)

type IncreKeeper struct {
	MaxValue    uint32 //Pool ve mat so ma minh dang co, gia toi da ma minh lay ra tu DB
	SeekerValue uint32 //Offset cua cai pool do, gia tri hien tai
}

func (self *IncreKeeper) FillValue(maxValue uint32, increStep uint32) {
	self.MaxValue = maxValue
	self.SeekerValue = maxValue - increStep
}

func (self *IncreKeeper) IsEmpty() bool {
	if self.SeekerValue >= self.MaxValue {
		return true
	}
	return false
}

func (self *IncreKeeper) GetNextValue() (uint32, bool) {
	if self.SeekerValue >= self.MaxValue {
		return 0, false
	}
	self.SeekerValue++
	return self.SeekerValue, true
}

var increKeeper IncreKeeper = IncreKeeper{}
var increStep uint32 = 100
var mutext *sync.Mutex = &sync.Mutex{}

// Cau hoi at ra, lam the nao de chuong trinh co the chay tot hon
// 1. Dam bao thoi gian <50ms
// 2. Khong bi dup du lieu
// 3. Co the chay dc tren nhieu may
func GetIncrementId(c *gin.Context, repo repo.SettingRepo) (gin.H, error) {
	mutext.Lock()
	if increKeeper.IsEmpty() {
		incre, _ := repo.GetNextId(increStep)
		increKeeper.FillValue(incre, increStep)
	}
	found, _ := increKeeper.GetNextValue()
	mutext.Unlock()
	return gin.H{
		"incre": found,
	}, nil
}
