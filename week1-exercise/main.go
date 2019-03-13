package main

import "fmt"
import "./helper"

func main() {
	// v := ""
	v := false
	// v := 0
	// v := nil
	fmt.Println("\n------------")
	if helper.IsEmpty(v) {
		fmt.Print("v is empty")
	} else {
		fmt.Print("v is empty")
	}
	fmt.Println("\n------------")
}
