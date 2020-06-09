package product

import (
	"../../models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

// Truyen db vao function
// GetProduct cua ban neu ban muon UT thi ban the nao?
// Neu ban de DB truc tiep vao thi lam sao mock dc BD
func (h ProductHandler) ProductGet(c *gin.Context) {
	// side effect
	product := models.Product{}
	id := c.Param("id")
	err := h.DB.Find(&product, id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, product)
}

func (h ProductHandler) Create(c *gin.Context) {
	p := &models.Product{}
	err := c.BindJSON(p) // Co the dung ca POST & GET
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := h.DB.Create(p)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, p)
}
