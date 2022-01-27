package main

import (
	"fmt"
	"github.com/dapo/webcrawler/crawler/implementation"
)

func main() {
	fmt.Println("Hello crawler")
	url := "https://gobyexample.com"
	downloadDir := "downloads/files"
	webCrawler := implementation.GetMyCrawler()
	webCrawler.Process(url, downloadDir)
}
