package handler

import (
	"strconv"

	"../helper"
	"../model"
	"../repo"

	"github.com/gin-gonic/gin"
)

func NoteCreate(c *gin.Context, noteRepo repo.NoteRepo) (*model.Note, error) {
	note := model.Note{}
	c.ShouldBind(&note)
	return noteRepo.Create(note)
}

func NoteGet(c *gin.Context, notePepo repo.NoteRepo) (*model.Note, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	return notePepo.Find(id)
}

func NoteList(c *gin.Context, notePepo repo.NoteRepo) ([]model.Note, error) {
	var pagination helper.Pagination
	c.ShouldBindQuery(&pagination)

	return notePepo.List(pagination)
}

func NoteUpdate(c *gin.Context, notePepo repo.NoteRepo) error {
	id, _ := strconv.Atoi(c.Param("id"))
	note := model.Note{}
	c.ShouldBind(&note)
	return notePepo.Update(id, note)
}

func NoteDelete(c *gin.Context, notePepo repo.NoteRepo) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return notePepo.Delete(id)
}
