package service

import (
	"context"
	"phudt/week4/internal/api"
	"phudt/week4/internal/model"
	"phudt/week4/internal/repo"
)

type service struct {
	patientRepo repo.Patient
	api.UnimplementedPatientServiceServer
}

func mutateAddRequestToModel(in *api.AddRequest) *model.Patient {
	p := &model.Patient{
		Fullname: in.Fullname,
		Address:  in.Address,
		Birthday: in.Birthday,
		Gender:   int32(in.Gender),
		Age:      in.Age,
	}
	return p
}

func mutateModelToAddAddResponse(m *model.Patient) *api.Patient {
	p := &api.Patient{
		Id:       m.Id,
		Fullname: m.Fullname,
		Address:  m.Address,
		Birthday: m.Birthday,
		Gender:   api.Gender(m.Gender),
		Age:      m.Age,
	}
	return p
}

func (s *service) Add(_ context.Context, in *api.AddRequest) (*api.Patient, error) {
	// TODO: Can suy nghi cho dat ten bien nay
	patient := mutateAddRequestToModel(in)
	p, err := s.patientRepo.Create(*patient)
	if err != nil {
		return nil, err
	}
	return mutateModelToAddAddResponse(p), nil
}

func NewService(repo repo.Patient) service {
	srv := service{patientRepo: repo}
	return srv
}
