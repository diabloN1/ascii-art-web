package myfunctions

// Merges the slices of string into a single string, while separating them with <br> tag.
func String(result []string) string {
	str := ""
	for _, v := range result {
		v = ReplaceSpaces(v)
		str += v + "<br>"
	}
	return str
}
