package main

import (
	"fmt"
)

func SquareArea(strArr []string, x0, y0 int) (area int) {
	if strArr[y0][x0] != '1' {
		return 0
	}
	maxY := len(strArr)
	maxX := len(strArr[0])
	square := 2
	for ; square <= maxX-x0 && square <= maxY-y0; square++ {
		if !isValidRow(strArr, square, x0, y0) {
			break
		}
		if !isValidCol(strArr, square, x0, y0) {
			break
		}
	}
	return (square - 1) * (square - 1)
}

func isValidRow(strArr []string, square, x0, y0 int) bool {
	for x := x0; x < x0+square; x++ {
		if strArr[y0+square-1][x] != '1' {
			return false
		}
	}
	return true
}

func isValidCol(strArr []string, square, x0, y0 int) bool {
	for y := y0; y < y0+square; y++ {
		if strArr[y][x0+square-1] != '1' {
			return false
		}
	}
	return true
}

func MaximalSquare(strArr []string) int {
	// Step 1 -> Visit to each point of matrix
	// Step 2 - Find max square at this point
	maxArea := 0
	maxY := len(strArr)
	maxX := len(strArr[0])
	for x := 0; x < maxX && maxX-x > maxArea; x++ {
		for y := 0; y < maxY && maxY-y > maxArea; y++ {
			area := SquareArea(strArr, x, y)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func main() {

	// do not modify below here, readline is our function
	// that operly reads in the input for you
	fmt.Println(MaximalSquare(readline()))
}

func readline() []string {
	return []string{"1111", "1101", "1111", "0111"}
}
