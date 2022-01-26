package crawler

import (
	"log"
	"os"
)

func createDirectory(path string) (err error) {
	if err = os.Mkdir(path, os.ModePerm); err != nil {
		log.Printf("Error occurred while creating directory: %s. ", err)
	}
	return
}
