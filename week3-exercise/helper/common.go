package helper

type Pagination struct {
	Page  uint `form:"p"`
	Limit uint `form:"l"`
}

func (self *Pagination) GetPage() uint {
	if self.Page == 0 {
		return 1
	}
	if self.Page > 100 {
		return 100
	}
	return self.Page
}

func (self *Pagination) GetOffset() uint {
	page := self.GetPage()
	limit := self.GetLimit()
	offset := (page - 1) * limit
	return offset
}

func (self *Pagination) GetLimit() uint {
	if self.Limit == 0 || self.Limit > 5 {
		return 5
	}
	return self.Limit
}
