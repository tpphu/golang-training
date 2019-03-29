package handler

import (
	"net/http"
	"strconv"

	"../helper"

	"github.com/gin-gonic/gin"
)

func NoteCreate(c *gin.Context) {
	var note Note
	c.ShouldBind(&note)
	noteRepo := NoteRepoImpl{}
	noteRepo.Create(note)
	c.JSON(200, note)
}

func NoteGet(c *gin.Context) {
	id := c.Param("id")
	var note Note
	err := db.Where("id = ?", id).First(&note).Error
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, note)
}

func NoteList(c *gin.Context) {
	var pagination helper.Pagination
	c.ShouldBindQuery(&pagination)
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	var notes []Note
	db.Offset(offset).
		Limit(limit).
		Find(&notes)
	c.JSON(http.StatusOK, notes)
}

func NoteUpdate(c *gin.Context) {
	// Xu ly du lieu tu client
	id, _ := strconv.Atoi(c.Param("id"))
	var note Note
	db.Where("id = ?", id).First(&note)
	c.ShouldBind(&note)
	note.ID = uint(id)
	// Update vo Db
	err := db.Model(&note).Updates(&note).Error
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	// Tra ket qua ve client
	c.JSON(http.StatusOK, note)
}

func NoteDelete(c *gin.Context) {
	// Xu ly du lieu tu client
	id, _ := strconv.Atoi(c.Param("id"))

	err := db.Where("id = ?", id).Delete(&Note{}).Error
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	// Tra ket qua ve client
	c.JSON(http.StatusOK, true)
}
