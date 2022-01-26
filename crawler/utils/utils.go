package utils

import (
	"golang.org/x/net/html"
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

func getLinks(body io.Reader) []string {
	var links []string
	tokenizers := html.NewTokenizer(body)
	for {
		tokenizer := tokenizers.Next()
		switch tokenizer {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := tokenizers.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
