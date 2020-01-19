package main

import (
	"flag"
	"log"

	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/gan-of-culture/go-hentai-scraper/config"
	"github.com/gan-of-culture/go-hentai-scraper/downloader"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/danbooru"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/nhentai"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/rule34"
	"github.com/gan-of-culture/go-hentai-scraper/static"
)

func init() {
	flag.StringVar(&config.Pages, "p", "", "Enter pages like 1,2,3-4,6,7,8-9")
	flag.StringVar(&config.OutputName, "o", "", "Output name")
	flag.StringVar(&config.OutputPath, "O", "", "Output path")
}

func download(url string) {
	var err error
	var data []static.Data
	domain := domainutil.Domain(url)
	switch domain {
	case "nhentai":
		data, err = nhentai.Extract(url)
	case "rule34":
		data, err = rule34.Extract(url)
	case "danbooru":
		data, err = danbooru.Extract(url)
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data {
		downloader.Download(d)
	}

}

func main() {
	flag.Parse()
	args := flag.Args()
	download(args[0])
}
