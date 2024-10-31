package myfunctions

// Returns a sliced string taking trailling spaces.
func ReplaceSpaces(str string) string {
	res := ""
	for i := range str {
		if OnlySpaces(str[i:]) {
			break
		}
		res += string(str[i])
	}
	return res
}
