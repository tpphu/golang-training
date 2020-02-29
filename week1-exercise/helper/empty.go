package helper

//
// TDD: return false
func IsEmpty(in interface{}) bool {
	if in == nil {
		return true
	}
	if in == 0 {
		return true
	}
	return false
}
