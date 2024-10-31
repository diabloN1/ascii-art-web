package myfunctions

// Returns bool, true in case to other char than spaces exists in the string.
func OnlySpaces(str string) bool {
	for _, v := range str {
		if v != ' ' {
			return false
		}
	}
	return true
}
