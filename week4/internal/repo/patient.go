package repo

import (
	"phudt/week4/internal/model"

	"gorm.io/gorm"
)

type patient struct {
	db *gorm.DB
}

type Filter struct {
	Field  string
	Method string
	Value  string
}

type Sort struct {
	Field string
	IsASC bool
}

type Pagination struct {
	Page  int
	Limit int
}

type Patient interface {
	Create(model.Patient) (*model.Patient, error)
	List(q string, filters []Filter, sort Sort, page Pagination) ([]model.Patient, error)
}

func NewPatientRepo(db *gorm.DB) Patient {
	return &patient{db: db}
}

// Code => Viet UT
func (r *patient) Create(m model.Patient) (*model.Patient, error) {
	result := &m
	err := r.db.Create(result).Error
	return result, err
	// return nil, nil
}

// Can phai viet UT cho cai nay
func (r *patient) List(q string, filters []Filter, sort Sort, page Pagination) ([]model.Patient, error) {
	return nil, nil
}
