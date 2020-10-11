package helper

// IndexOf vs indexOf
func IndexOf(arr []int, find int) int {
	for i, v := range arr {
		if v == find {
			return i
		}
	}
	return -1
}

// IndexOfString to search the position of a string in a slice of string.
func IndexOfString(arr []string, find string) int {
	for i, v := range arr {
		if v == find {
			return i
		}
	}
	return -1
}
