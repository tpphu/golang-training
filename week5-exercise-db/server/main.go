package main

import (
	"fmt"

	"context"

	model "../model"
	proto "../proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro"
)

type serviceImpl struct {
	DB *gorm.DB
}

func (self *serviceImpl) Create(ctx context.Context, req *proto.NoteCreateReq, res *proto.Note) error {
	note := model.Note{
		Title:     req.Title,
		Completed: req.Completed,
	}
	err := self.DB.Create(&note).Error
	if err != nil {
		return err
	}
	res.Id = uint32(note.ID)
	res.Title = note.Title
	res.Completed = note.Completed
	return nil
}

func main() {
	// 1. Create service
	service := micro.NewService(
		micro.Name("note-service"), // Quan trong la co mot cai ten
	)
	// 1.1 Registry to Consul
	service.Init()

	// 2. Register handler
	db, _ := gorm.Open("mysql", "default:secret@/notes?charset=utf8&parseTime=True&loc=Local")
	serviceImpl := serviceImpl{
		DB: db,
	}
	proto.RegisterNoteServiceHandler(service.Server(), &serviceImpl)

	// 3.  Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
