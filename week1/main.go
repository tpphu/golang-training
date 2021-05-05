package main

import (
	"encoding/json"
	"fmt"
	"phudt/week1/helper"
	"phudt/week1/model"
	"strconv"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Sum of: {1, 2, 3, 4} is ", helper.Sum(1, 2, 3, 4))
	p := model.Person{
		Name:          "Phu",
		Year_of_birth: 1986,
	}
	data, _ := json.Marshal(p)
	fmt.Println("JSON marshal: ", string(data))

	s := "100a000"
	fmt.Println(" len of '%s' = ", len(s))
	var n int
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" Number of %s is %d\n", s, n)
}
