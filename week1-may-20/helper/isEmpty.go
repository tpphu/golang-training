package helper

// func isEmpty(input int) bool {
// 	if input == 0 {
// 		return true
// 	}
// 	return false
// }

func isEmpty(input interface{}) bool {
	if input == 0 {
		return true
	}
	if input == "" {
		return true
	}
	if input == false {
		return true
	}
	if input == nil {
		return true
	}
	return false
}