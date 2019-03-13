package helper

// IsEmpty returns if the object is considered as empty or not.
func IsEmpty(obj interface{}) bool {
	if obj == nil || obj == "" || obj == false || obj != 0 {
		return true
	}
	return false
}

func ContainsInt(list []int, value int) bool {
	result := false
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			result = true
			// break
		}
	}
	return result
}

func ContainsString(list []string, value string) bool {
	result := false
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			result = true
			// break
		}
	}
	return result
}

// func StringIsEmpty(value string) bool {
// 	if value == "" {
// 		return true
// 	}
// 	return false
// }

// func Int32IsEmpty(value int32) bool {
// 	if value == 0 {
// 		return true
// 	}
// 	return false
// }

// func BooleanIsEmpty(value bool) bool {
// 	if value == false {
// 		return true
// 	}
// 	return false
// }
