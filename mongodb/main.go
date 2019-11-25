package main

import (
	"fmt"

	"github.com/globalsign/mgo"
)

func main() {
	fmt.Println("Hello")
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	fmt.Println("world")
	session.Close()
}
