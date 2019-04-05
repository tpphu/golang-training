package handler

import (
	"../repo"
	"github.com/gin-gonic/gin"
)

var increStep uint32 = 1

// Cau hoi at ra, lam the nao de chuong trinh co the chay tot hon
// 1. Dam bao thoi gian <50ms
// 2. Khong bi dup du lieu
// 3. Co the chay dc tren nhieu may
func GetIncrementId(c *gin.Context, repo repo.SettingRepo) (gin.H, error) {
	incre, _ := repo.GetNextId(increStep)
	return gin.H{
		"incre": incre,
	}, nil
}
