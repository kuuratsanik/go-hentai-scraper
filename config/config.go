package config

var (
	// OutputPath
	OutputPath string
	// OutputName
	OutputName string
	// Pages of doujinshi
	Pages string
	// RestrictContent
	RestrictContent bool
	// SelectStream
	SelectStream string
	// ShowInfo of all avaiable streams
	ShowInfo bool
	// Username for exhentai.org
	Username string
	// UserPassword for exhentai.org
	UserPassword string
)

// FakeHeaders fake http headers
var FakeHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Language": "en-US,en;q=0.8",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
}
