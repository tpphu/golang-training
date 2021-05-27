package handler

import (
	"net/http"
	"phudt/week3/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Patient interface {
	PatientCreate(c *gin.Context)
}

type patient struct {
	db *gorm.DB
}

func NewPatient(db *gorm.DB) Patient {
	return patient{
		db: db,
	}
}

func (p patient) PatientCreate(c *gin.Context) {
	patient := model.Patient{}
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err := p.db.Create(&patient).Error
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, patient)
}
