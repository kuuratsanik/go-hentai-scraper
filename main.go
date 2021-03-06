package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gan-of-culture/go-hentai-scraper/config"
	"github.com/gan-of-culture/go-hentai-scraper/downloader"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/danbooru"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/ehentai"
	"github.com/gan-of-culture/go-hentai-scraper/utils"

	//"github.com/gan-of-culture/go-hentai-scraper/extractor/hanime"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/nhentai"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/rule34"
	"github.com/gan-of-culture/go-hentai-scraper/extractor/underhentai"
	"github.com/gan-of-culture/go-hentai-scraper/static"
)

func init() {
	flag.StringVar(&config.OutputName, "o", "", "Output name")
	flag.StringVar(&config.OutputPath, "O", "", "Output path")
	flag.StringVar(&config.Pages, "p", "", "Enter pages like 1,2,3-4,6,7,8-9 for doujins")
	flag.BoolVar(&config.RestrictContent, "r", false, "Don't scrape Restricted Content")
	flag.StringVar(&config.SelectStream, "s", "0", "Select a stream")
	flag.BoolVar(&config.ShowInfo, "i", false, "Show info")
	flag.StringVar(&config.Username, "un", "", "Username for exhentai/forum e hentai")
	flag.StringVar(&config.Username, "up", "", "User password for exhentai/forum e hentai")
}

func download(url string) {
	var err error
	var data []static.Data
	re := regexp.MustCompile("http?s://([^\\.]+)")
	matches := re.FindStringSubmatch(url)
	if len(matches) < 2 {
		log.Fatal("Can't parse URL")
	}
	if matches[1] == "www" {
		re := regexp.MustCompile("http?s://www.([^\\.]+)")
		matches = re.FindStringSubmatch(url)
	}

	switch matches[1] {
	case "danbooru":
		data, err = danbooru.Extract(url)
	case "e-hentai":
	case "exhentai":
		data, err = ehentai.Extract(url)
	//case "hanime":
	//data, err = hanime.Extract(url)
	case "nhentai":
		data, err = nhentai.Extract(url)
	case "rule34":
		data, err = rule34.Extract(url)
	case "underhentai":
		data, err = underhentai.Extract(url)
	default:
		data = append(data, static.Data{
			Site:  "unknown",
			Type:  "unknown",
			Title: "Universal download",
			Streams: map[string]static.Stream{
				"0": static.Stream{
					URLs: []static.URL{
						{
							URL: url,
							Ext: utils.GetLastItemString(strings.Split(url, ".")),
						},
					},
				},
			},
			Url: url,
		})
	}
	if err != nil {
		log.Fatal(err)
	}

	if config.ShowInfo {
		for _, singleData := range data {
			jsonData, _ := json.MarshalIndent(singleData, "", "    ")
			fmt.Printf("%s\n", jsonData)
		}
		return
	}

	for _, d := range data {
		err = downloader.Download(d)
		if err != nil {
			log.Println(err)
		}
	}

}

func main() {
	flag.Parse()
	args := flag.Args()
	download(args[0])
}
