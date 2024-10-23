package myfunctions

import (
	"log"
	"os"
)

func Read(fileName string) ([]byte, error) {
	
	//Open File.
    file, err := os.Open(fileName)

	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	defer file.Close()

	//Get file info.
    fileInfo, err := file.Stat()
    if err != nil {
        log.Println("Error getting file stats:", err)
		return []byte{}, err
    }

    //Get file size.
    fileSize := fileInfo.Size()
    data := make([]byte, fileSize)

	//Reading data.
    _, err = file.Read(data)
    if err != nil {
		log.Println("Error reading the file:", err)
		return []byte{}, err
    }
	return data, nil
}
