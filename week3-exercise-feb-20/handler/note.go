package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/week2-exercise/form"
	"github.com/tpphu/week2-exercise/helper"
	"github.com/tpphu/week2-exercise/model"
)

// Dung de mapping data tu nguoi dung goi len
// Dung de lam validation
type NoteHandler struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

func (h NoteHandler) inject() {
	h.Engine.POST("/note", h.create)
	h.Engine.GET("/note/:id", h.get) // matching
	h.Engine.GET("/note", h.list)
	h.Engine.PUT("/note/:id", h.update)
	h.Engine.DELETE("/note/:id", h.delete)
}

// O day chung ta se giai quyet bai toan phan tran gthe nao
func (h NoteHandler) list(c *gin.Context) {
	pager, err := helper.NewPaginationFromRequest(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	notes := []model.Note{}
	// Code cho nay co van de
	// Offset va Limit doi voi data lon
	err = h.DB.
		Limit(pager.GetLimit()).
		Offset(pager.GetOffset()).
		Find(&notes).
		Error
	if err != nil {
		c.JSON(404, gin.H{
			"code":    2000,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, notes)
	return
}

func (h NoteHandler) delete(c *gin.Context) {
	// Lay doi tuong muon update
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw) //Ascii
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	// Delete
	err = h.DB.Delete(model.Note{}, "id = ?", id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Deleted is success",
	})
}

func (h NoteHandler) update(c *gin.Context) {
	// Lay doi tuong muon update
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw) //Ascii
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	// Lay du lieu tu client goi len va validate no
	input := form.Note{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	// Update note
	note := model.Note{}
	note.ID = uint(id)
	// https://github.com/jinzhu/gorm/issues/202
	// https://github.com/jinzhu/gorm/issues/2119
	err = h.DB.Model(&note).
		Update("title", input.Title).
		Update("completed", input.Completed).
		Error
	if err != nil {
		c.JSON(404, gin.H{
			"code":    2000,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, note)
}

func (h NoteHandler) get(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw) //Ascii
	if err != nil {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	note := model.Note{}
	err = h.DB.First(&note, id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"code":    2000,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, note)
}

func (h NoteHandler) create(c *gin.Context) {
	// Lay du lieu tu client goi len va validate no
	input := form.Note{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	// Fill du lieu tu nguoi dung vao model
	note := model.Note{}
	note.Fill(input)
	// Insert vao db
	err := h.DB.Create(&note).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":    2000,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, note)
}
