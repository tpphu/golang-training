package service

import (
	"context"
	"errors"
	"phudt/week4/internal/api"
	"phudt/week4/internal/model"
	"phudt/week4/internal/repo"
	"testing"
)

/**
|-------------------------------------------------------------------------
| Ban chat mock -> viet lai cac struct xu ly cua mot interface de phu ho
| voi muc dich test test.
|-----------------------------------------------------------------------*/

type mockRepo struct {
}

// Khac nao em dang lap trinh ra mot thu khac
func (mock *mockRepo) Create(m model.Patient) (*model.Patient, error) {
	if m.Age < 0 {
		return nil, errors.New("Age should not be 0")
	}
	return &model.Patient{
		Id:       1,
		Fullname: "",
		Address:  "",
		Birthday: "",
		Gender:   0,
		Age:      0,
	}, nil
}

func (r *mockRepo) List(q string, filters []repo.Filter, sort repo.Sort, page repo.Pagination) ([]model.Patient, error) {
	return nil, nil
}

func TestAdd_HappyCase(t *testing.T) {
	patientRepo := &mockRepo{}
	s := service{
		patientRepo: patientRepo,
	}
	in := &api.AddRequest{
		Fullname: "Phu",
		Address:  "HCM",
		Birthday: "2021-10-10",
		Gender:   0,
		Age:      0,
	}
	out, err := s.Add(context.Background(), in)
	if err != nil {
		t.Error(err)
	}
	if out.GetId() == 0 {
		t.Error("Id should not be 0")
	}
}

func TestAdd_TestError(t *testing.T) {
	patientRepo := &mockRepo{}
	s := service{
		patientRepo: patientRepo,
	}
	in := &api.AddRequest{
		Fullname: "Phu",
		Address:  "HCM",
		Birthday: "2021-10-10",
		Gender:   0,
		Age:      -1,
	}
	_, err := s.Add(context.Background(), in)
	if err == nil {
		t.Error("Should not be nil")
	}
}
