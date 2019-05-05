package main

import (
	"fmt"
	"reflect"
)

type Note struct {
	Title     string `binding:"required,min=6,max=255"`
	Completed bool
}

func (self Note) GetTitle() string {
	return self.Title
}
func (self *Note) SetTitle(title string) {
	self.Title = title
}

func main() {
	note := &Note{}
	typeOfVar := reflect.TypeOf(note)
	if typeOfVar.Kind() == reflect.Ptr {
		fmt.Println("Name: ", typeOfVar.Elem().Name())
	} else {
		fmt.Println("Name: ", typeOfVar.Name())
	}

	fmt.Println("Kind: ", typeOfVar.Kind())
	if typeOfVar.Kind() == reflect.Ptr {
		fmt.Println("NumMethod: ", typeOfVar.Elem().NumMethod())
	} else {
		fmt.Println("NumMethod: ", typeOfVar.NumMethod())
	}

	valueOfVar := reflect.ValueOf(note).Elem()
	fieldOfTitle := valueOfVar.FieldByName("Title")
	fieldOfTitle.SetString("Todo 2")

	fmt.Println("note: ", note)
}
