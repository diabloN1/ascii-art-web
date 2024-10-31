package asciiArtFs

import (
	"asciiArtWeb/asciiArtFs/myFunctions"
	"log"
	"strings"
)

func AsciiArtFs(text string, banner string) (string, any) {
	banner = "asciiArtFs/banners/" + banner + ".txt"
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
	res := myfunctions.String(result)
	return res, nil
}
