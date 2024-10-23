package asciiArtFs

import (
	"asciiArtWeb/asciiArtFs/myFunctions"
	"log"
	"strings"
)

func AsciiArtFs(text string, banner string) (string, any) {
	banner = "asciiArtFs/" + banner + ".txt"
	standard, err := myfunctions.Read(banner)
	if err != nil {
		switch {
		case strings.HasSuffix(err.Error(), "no such file or directory"):
			return "", "NotFound"
		default :
			return "", ""
		}
	}
	asciiChars := myfunctions.BytesToAsciiMap(standard)
	result, err := myfunctions.WriteResult(text, asciiChars)
	if err != nil {
		log.Println(err)
		return "", "Non-Ascii"
	}	
	res := String(result)
	return res, nil
}

func String(result []string) string {
	str := ""
	for _, v := range result {
		v = replaceSpaces(v)
		str += v + "<br>"
	}
	return str
}

func replaceSpaces(str string) string {
	res := ""
	for i := range str {
		if OnlySpaces(str[i:]) {
			break
		}
		res += string(str[i])
	}
	return res
}

func OnlySpaces(str string) bool {
	for _, v := range str {
		if v != ' ' {
			return false
		}
	}
	return true
}