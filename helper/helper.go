package helper

func IsLengthValid(length, min, max int) bool {
	if length >= min && length <= max {
		return true
	} else {
		return false
	}
}
