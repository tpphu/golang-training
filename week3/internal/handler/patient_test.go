package handler

import (
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestPatientCreate(t *testing.T) {
	// dsn := "root:root@tcp(127.0.0.1:3307)/covid?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	t.Error(err)
	// }

	//mock db
	dbMock, _, _ := sqlmock.New()
	dialect := mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       "",
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	})
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	p := patient{
		db: db,
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	p.PatientCreate(c)
	if w.Code != 200 {
		t.Errorf("HTTP Status Code should be 200, but got : %d", w.Code)
	}
}
