package helper

func Sum(arr ...int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}
