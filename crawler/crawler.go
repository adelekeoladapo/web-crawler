package crawler

type Crawler interface {
	Process(url, dir string)
}
