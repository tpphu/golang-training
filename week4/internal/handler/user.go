package handler

import (
	"phudt/week4/internal/model"
	"phudt/week4/internal/provider"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	rp *provider.Provider
}

func NewUserHandler(rp *provider.Provider) UserHandler {
	return UserHandler{
		rp: rp,
	}
}

func (h UserHandler) Create(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, err)
		return
	}
	h.rp.DB.Create(&user)
	c.JSON(200, user)
}

func (h UserHandler) CreateWithDepartment(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, err)
		return
	}
	h.rp.DB.Create(&user)
	c.JSON(200, user)
}
