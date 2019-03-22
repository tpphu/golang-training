package main

import "fmt"

func SquareArea(strArr []string, x0, y0 int) (area int) {
	rows := len(strArr)
	cols := len(strArr[0])
	if strArr[x0][y0] != '1' {
		return 0
	}
	vertex := 1
	isStop := false
	for ; ; vertex++ {
		if vertex >= (cols-y0) || vertex >= (rows-x0) {
			break
		}
		vRows := x0 + vertex
		vCols := y0 + vertex
		// Theo dong
		for i := y0; i < vCols; i++ {			
			if strArr[vRows][i] != '1' {
				isStop = true
				break
			}
		}
		// Theo cot
		for i := x0; i < vRows; i++ {
			if strArr[i][vCols] != '1' {
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

func readline() []string{
	return []string{"10100", "10111", "11111", "10010"}
}
