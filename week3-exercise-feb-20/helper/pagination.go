package helper

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var MAX_LIMIT uint = 50
var MAX_PAGE uint = 100

// Chung ta giai quyet cai bai toan phan trang nhu the nao
type Pagination struct {
	Page  uint
	Limit uint
}

func NewPaginationFromRequest(c *gin.Context) (Pagination, error) {
	pagination := Pagination{}
	pageRaw := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageRaw)
	if err != nil {
		return pagination, err
	}
	if page <= 0 {
		return pagination, errors.New("Page should be larger than 0")
	}
	if uint(page) > MAX_PAGE {
		return pagination, errors.New("Page should not be larger than max limit")
	}
	pagination.Page = uint(page)

	limitRaw := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitRaw)
	if err != nil {
		return pagination, err
	}
	if limit <= 0 {
		return pagination, errors.New("Limit should be larger than 0")
	}
	if uint(limit) > MAX_LIMIT {
		return pagination, errors.New("Limit should not be larger than max limit")
	}
	pagination.Limit = uint(limit)
	return pagination, nil
}

// Offset co loi gi
func (p Pagination) GetOffset() uint {
	return (p.GetPage() - 1) * p.Limit
}

// Han che cac loi ve bang thong, bandwidth
func (p Pagination) GetLimit() uint {
	if p.Limit > 50 {
		return MAX_LIMIT
	}
	return p.Limit
}

//
func (p Pagination) GetPage() uint {
	if p.Page > MAX_PAGE {
		return MAX_PAGE
	}
	return p.Page
}
