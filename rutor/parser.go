package rutor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/bamboV/torrent"
	"strconv"
	"strings"
)

type Parser struct {
	BaseURL string
}

func (p *Parser) Parse(searchPhrase string) ([]torrent.DistributionLine, error) {
	url := p.BaseURL + "/search/" + searchPhrase

	lines := []torrent.DistributionLine{}

	document, err := goquery.NewDocument(url)

	if err != nil {
		return nil, err
	}

	document.Find("div#content").Find("div#index").Find("table").Find("tr").Each(func(index int, selection *goquery.Selection) {
		class, exists := selection.Attr("class")
		if exists && class == "backgr" {
			return
		}
		col := selection.Find("td")
		line := torrent.DistributionLine{}
		secondCol := col.First().Next()
		secondCol.Find("a").Each(func(i int, link *goquery.Selection) {
			switch i {
			case 0:
				line.DownloadLink, _ = link.Attr("href")
				line.Id, _ = strconv.Atoi(strings.Trim(line.DownloadLink, "/download/"))
				break
			case 1:
				line.MagnetLink, _ = link.Attr("href")
				break
			case 2:
				line.TopicLink, _ = link.Attr("href")
				line.Title = link.Text()
				break
			}
		})

		line.Size = secondCol.Next().Next().Text()

		lines = append(lines, line)
	})

	return lines, nil
}
