package main

import (
	"flag"
	"fmt"
	"github.com/dapo/webcrawler/crawler/implementation"
)

func main() {
	fmt.Println("Hey Simple Crawler")
	url := flag.String("url", "https://gobyexample.com", "Starting URL")
	downloadDir := flag.String("dir", "downloads/files", "Download directory")
	flag.Parse()
	webCrawler := implementation.GetMyCrawler()
	webCrawler.Process(*url, *downloadDir)
}
