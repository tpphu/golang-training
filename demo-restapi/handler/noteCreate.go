package handler

import (
	"strconv"

	"../config"
	"../model"
	"../repo"
	"github.com/gin-gonic/gin"
)

func PreProcessingNoteInput(c *gin.Context) (model.Note, error) {
	input := model.Note{}
	err := c.BindJSON(&input)
	if err != nil {
		return input, err
	}
	identity, _ := c.Get(config.IdentityKey)
	authorId, _ := strconv.Atoi(identity.(string))
	input.AuthorID = uint(authorId)
	return input, nil
}

func CreateNoteHandler(noteRepo repo.NoteRepo, input model.Note) (model.Note, error) {
	note, err := noteRepo.Create(input)
	return note, err
}
