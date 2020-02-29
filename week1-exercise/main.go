package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Welcome to Golang Course from Phu Phong Tran")

	understandSlice()
	fmt.Println("---")
	understandBuildSlice()
	fmt.Println("---")
	understandStruct()
}

func understandSlice() {
	arr := []int{1, 2, 3, 4}

	typeOfArr := reflect.TypeOf(arr)
	fmt.Println("Name of typeOfArr:", typeOfArr.String())

	kindOfArr := typeOfArr.Kind()
	fmt.Println("Name of kindOfArr:", kindOfArr.String())

	kindOfElem := typeOfArr.Elem().Kind()
	fmt.Println("Name of kindOfElem:", kindOfElem.String())

	valueOfArr := reflect.ValueOf(arr)
	value := valueOfArr.Index(0).String()
	fmt.Println("#1 Value of valueOfArr:", value)

}

func understandBuildSlice() {
	slice := reflect.MakeSlice(reflect.TypeOf([]int{}), 3, 3)
	slice.Index(0).Set(reflect.ValueOf(1))
	slice.Index(1).Set(reflect.ValueOf(2))
	slice.Index(2).Set(reflect.ValueOf(3))
	fmt.Println("Value of slice: ", slice)
}
func understandStruct() {
	type Note struct {
		ID    int
		Title string
	}
	note := Note{}
	typeOfNote := reflect.TypeOf(note)
	fmt.Println("Name of typeOfNote:", typeOfNote.String())

	kindOfNote := typeOfNote.Kind()
	fmt.Println("Name of typeOfNote:", kindOfNote.String())
}
