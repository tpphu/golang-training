package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	pb "github.com/tpphu/week4-exercise/proto"
)

type Note struct {
	gorm.Model
	Title     string
	Completed bool
}

func (n Note) Populate(note *pb.Note) error {
	note.Id = int32(n.ID)
	note.Title = n.Title
	note.Completed = n.Completed
	note.CreatedAt = &timestamp.Timestamp{Seconds: n.CreatedAt.Unix()}
	note.UpdatedAt = &timestamp.Timestamp{Seconds: n.UpdatedAt.Unix()}
	return nil
}
