package main

import "fmt"

func SquareArea(strArr []string, x0, y0 int) (area int) {
	rows := len(strArr)
	cols := len(strArr[0])
	vertex := 0
	isStop := false
	for ; ; vertex++ {
		if vertex >= (rows-y0) || vertex >= (cols-x0) {
			break
		}
		for y := y0 + vertex; y < cols; y++ {
			if strArr[y0+vertex][y] != '1' {
				isStop = true
				break
			}
		}
		for x := x0 + vertex; x < rows; x++ {
			if strArr[x][x0+vertex] != '1' {
				isStop = true
				break
			}
		}
		if isStop {
			break
		}
	}
	return vertex * vertex
}

/**
 * Muc tieu cua func nay la de:
 * 1. Tim cac squre di tu vi tri [x, y]
 * 2. So sanh voi square truoc do de cap nhat square lon hon
 */
func MaximalSquare(strArr []string) (maxArea int) {
	rows := len(strArr)
	cols := len(strArr[0])
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
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
	// that properly reads in the input for you
	fmt.Println(MaximalSquare(readline()))

}

func readline() []string {
	// return []string{"0111", "1111", "1111", "1111"}
	return []string{"10100", "10111", "11111", "10010"}
}
