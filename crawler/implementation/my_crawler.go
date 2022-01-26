package implementation

import (
	"github.com/dapo/webcrawler/crawler"
)

type MyCrawler struct {
}

func (o *MyCrawler) Process(url, dir string) {

}

func GetMyCrawler() crawler.Crawler {
	return &MyCrawler{}
}
