package implementation

import (
	"errors"
	"fmt"
	"github.com/dapo/webcrawler/crawler"
	"github.com/dapo/webcrawler/crawler/utils"
	"github.com/dapo/webcrawler/crawler/utils/queue"
	"log"
	"path"
	"strings"
)

type MyCrawler struct {
	pendingUrls queue.SimpleQueue
	crawledUrls []string
}

func (o *MyCrawler) Process(startingUrl, downloadDir string) {
	if err := utils.CreateDirectory(downloadDir); err != nil {
		log.Fatal("Error occurred while creating directory ", err)
	}
	if strings.HasSuffix(startingUrl, "/") {
		startingUrl = startingUrl[:len(startingUrl)-1]
	}
	o.pendingUrls.Enqueue(startingUrl)
	for !o.pendingUrls.Empty() {
		url, err := o.pendingUrls.Dequeue()
		if err != nil {
			log.Println("Error occurred while fetching from queue. ", err)
		}
		strUrl := url.(string)
		if o.isUrlCrawled(strUrl) {
			continue
		}
		fmt.Println(" . . .  ", url, " . . .")
		body, err := utils.LoadUrl(strUrl)
		if err != nil {
			log.Println("Error occurred while loading url ", strUrl)
		}
		filename := path.Base(strUrl)
		err = utils.DownloadWebPage(downloadDir, filename, body)
		body, _ = utils.LoadUrl(strUrl)
		urls := utils.GetLinks(body)
		o.updatePendingUrls(strUrl, urls)
		o.crawledUrls = append(o.crawledUrls, strUrl)
	}
	fmt.Println(len(o.crawledUrls), " urls crawled successfully")
}

func (o *MyCrawler) isUrlCrawled(url string) bool {
	for _, u := range o.crawledUrls {
		if u == url {
			return true
		}
	}
	return false
}

func (o *MyCrawler) updatePendingUrls(baseUrl string, urls []string) {
	for _, u := range urls {
		goodUrl, err := prepareFullUrl(baseUrl, u)
		if err == nil {
			o.pendingUrls.Enqueue(goodUrl)
		}
	}
}

func prepareFullUrl(baseUrl, url string) (goodUrl string, err error) {
	if (strings.HasPrefix(url, "www") || strings.HasPrefix(url, "http")) && !strings.HasPrefix(url, baseUrl) {
		err = errors.New("bad url")
	} else if (strings.HasPrefix(url, baseUrl)) && baseUrl != url {
		goodUrl = url
	} else if strings.HasPrefix(url, "/") && strings.HasSuffix(url, "/") {
		goodUrl = baseUrl + url[:len(url)-1]
	} else if strings.HasPrefix(url, "/") && !strings.HasSuffix(url, "/") {
		goodUrl = baseUrl + url
	} else if !strings.HasPrefix(url, "/") && !strings.HasSuffix(url, "/") {
		goodUrl = baseUrl + "/" + url
	} else if !strings.HasPrefix(url, "/") && strings.HasSuffix(url, "/") {
		goodUrl = baseUrl + "/" + url[:len(url)-1]
	} else {
		err = errors.New("bad url")
	}
	return
}

func GetMyCrawler() crawler.Crawler {
	return &MyCrawler{}
}
