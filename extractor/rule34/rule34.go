package rule34

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gan-of-culture/go-hentai-scraper/request"
	"github.com/gan-of-culture/go-hentai-scraper/static"
	"github.com/gan-of-culture/go-hentai-scraper/utils"
)

const site = "https://rule34.paheal.net"

// Extract for data
func Extract(URL string) ([]static.Data, error) {

	data := []static.Data{}

	for _, element := range ParseURL(URL) {

		htmlString, err := request.Get(site + element)
		if err != nil {
			return nil, err
		}

		elementTag, err := extractData(htmlString)
		if err != nil {
			return nil, err
		}

		size, err := request.Size(elementTag["src"], site+element)
		if err != nil {
			return nil, errors.New("[Danbooru]No image size not found")
		}

		streams := make(map[string]static.Stream)
		streams["0"] = static.Stream{
			URLs: []static.URL{
				{
					URL: elementTag["src"],
					Ext: utils.GetLastItemString(strings.Split(elementTag["src"], ".")),
				},
			},
			Quality: fmt.Sprintf("%s x %s", elementTag["data-width"], elementTag["data-height"]),
			Size:    size,
		}

		data = append(data, static.Data{
			Site:    site,
			Title:   elementTag["title"],
			Type:    elementTag["type"],
			Streams: streams,
			Url:     URL,
		})

	}

	return data, nil
}

// ParseURL data
func ParseURL(URL string) []string {
	htmlString, err := request.Get(URL)
	if err != nil {
		return nil
	}

	doc := soup.HTMLParse(htmlString)
	items := doc.FindAll("a", "class", "shm-thumb-link")

	// overview page | get URL to all elements
	content := make([]string, len(items))
	for idx, item := range items {
		content[idx] = item.Attrs()["href"]
	}

	if len(content) != 0 {
		return content
	}

	re := regexp.MustCompile("[0-9]{6,8}")
	id := re.FindString(URL)
	if id == "" {
		return nil
	}

	content = []string{"/post/view/" + id}

	return content
}

func extractData(htmlString string) (map[string]string, error) {
	doc := soup.HTMLParse(htmlString)
	mainTag := doc.Find("img", "class", "shm-main-image")
	if mainTag.Error != nil {
		mainTag = doc.Find("video", "class", "shm-main-image")
		if mainTag.Error != nil {
			return nil, mainTag.Error
		}

	}

	attrs := mainTag.Attrs()
	attrs["title"] = doc.Find("input", "name", "tag_edit__tags").Attrs()["value"]

	attrs["type"] = "image"
	if strings.Contains(attrs["src"], ".gif") {
		attrs["type"] = "gif"
	}

	if attrs["data-width"] != "" {
		return attrs, nil
	}

	attrs["type"] = "video"

	// get the src attr of the source tag
	attrs["src"] = doc.Find("section", "id", "Videomain").FindAll("a")[0].Attrs()["href"]
	re := regexp.MustCompile("[a-z]+:[\t\n\f\r ][0-9]+px")
	dimensions := re.FindAllString(attrs["style"], -1)

	for _, dimension := range dimensions {
		splitKey := strings.Split(dimension, ": ")

		// splitKey[0] = width/height splitKey[1] = ?px

		switch splitKey[0] {
		case "width":
			attrs["data-width"] = splitKey[1]
		case "height":
			attrs["data-width"] = splitKey[1]
		default:
			return nil, errors.New("[Rule34]Can't calc video size")
		}
	}

	return attrs, nil
}
