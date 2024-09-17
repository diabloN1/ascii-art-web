package asciiArtFs

import (
	"fmt"
	"log"
	"asciiArtWeb/asciiArtFs/myFunctions"
)

func AsciiArtFs(text string, banner string) (string, error) {
	banner = "asciiArtFs/" + banner + ".txt"
	standard, err := myfunctions.Read(banner)
	if err != nil {
		return "", fmt.Errorf("")
	}
	asciiChars := myfunctions.BytesToAsciiMap(standard)
	result, err := myfunctions.WriteResult(text, asciiChars)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("")
	}
	res := String(result)
	return res, nil
}

func String(result []string) string {
	str := ""
	for _, v := range result {
		str += v + "\n"
	}
	return "\n" + str
}