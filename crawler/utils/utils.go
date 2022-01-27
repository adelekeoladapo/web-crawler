package utils

import (
	"errors"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func CreateDirectory(path string) (err error) {
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		log.Printf("Error occurred while creating directory: %s. ", err)
	}
	return
}

func DownloadWebPage(dir string, filename string, body io.Reader) (err error) {
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

func GetLinks(body io.Reader) []string {
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

func LoadUrl(url string) (body io.Reader, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error occurred while loading url ", url)
		err = errors.New("Error occurred while loading url ")
		return
	}
	body = resp.Body
	return
}
