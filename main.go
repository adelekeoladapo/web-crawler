package main

import (
	"fmt"
	"github.com/dapo/webcrawler/crawler/implementation"
)

func main() {
	fmt.Println("Hello crawler")
	url := "google.com"
	downloadDir := "downloads/files"

	webCrawler := implementation.GetMyCrawler()

	webCrawler.Process(url, downloadDir)

}
