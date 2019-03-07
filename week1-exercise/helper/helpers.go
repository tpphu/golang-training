package helper

// IsEmpty returns if the object is considered as empty or not.
func IsEmpty(obj interface{}) bool {
	if obj == nil || obj == "" || obj == false || obj == 0 {
		return true
	}

	return false
}
