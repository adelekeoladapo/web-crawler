package crawler

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func createDirectory(path string) (err error) {
	if err = os.Mkdir(path, os.ModePerm); err != nil {
		log.Printf("Error occurred while creating directory: %s. ", err)
	}
	return
}

func downloadWebPage(dir string, filename string, body io.Reader) (err error) {
	file, err := os.Create(filepath.Join(dir, filepath.Base(filename+".html")))
	if err != nil {
		log.Println("Error occurred while creating file. ", err)
		return
	}
	_, err = io.Copy(file, body)
	if err != nil {
		log.Println("Error occurred while downloading file. ", err)
		return
	}
	return
}
