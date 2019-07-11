package main

import "fmt"
import "unicode/utf8"

func main() {
	showDebug5()
}

func showDebug1() {
	var s string
	s = "Xin chao"
	fmt.Println("Length of string: ", len(s))
	fmt.Println(s)

	bytes := []byte(s)
	fmt.Println("Length of bytes: ", len(bytes))
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%v ", bytes[i])
	}
	fmt.Println("")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%x ", bytes[i])
	}
}

func showDebug2() {
	var s string
	s = "Xin chào"
	fmt.Println("Length of string: ", len(s))
	fmt.Println(s)

	bytes := []byte(s) //UTF8 encoding
	fmt.Println("Length of bytes: ", len(bytes))
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%v ", bytes[i])
	}
	fmt.Println("")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%x ", bytes[i])
	}
}

func showDebug3() {
	var s string
	s = "Xin chào"
	fmt.Println("Length of string: ", len(s))
	fmt.Println(s)

	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	bytes := []byte(s)
	fmt.Println("Length of bytes: ", len(bytes))
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%v ", bytes[i])
	}
	fmt.Println("")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%x ", bytes[i])
	}
}

func showDebug4() {
	// https://golang.org/doc/go1#rune
	// Rune is a Type.
	// It occupies 32bit and is meant to represent a Unicode CodePoint.
	var s string
	s = "Xin chào"
	fmt.Println("Length of string: ", utf8.RuneCountInString(s))
	fmt.Println(s)

	fmt.Println("")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// fmt.Printf("%v ", runes[i])
		// fmt.Printf("%c ", runes[i])
		fmt.Printf("%x ", runes[i])
	}
	fmt.Println("")
}

func showDebug5() {
	s := "Xin chào"

	for index, char := range s {
		fmt.Printf("character at %d is %c\n", index, char)
	}
	fmt.Println("")
}
