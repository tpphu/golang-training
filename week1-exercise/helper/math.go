package helper

func Max(list []int) int {
	max := list[0]
	for i := 0; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
		}
	}
	return max
}
