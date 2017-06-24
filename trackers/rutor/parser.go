package rutor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/bamboV/torrent"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	BaseURL string
}

func (p *Parser) Parse(searchPhrase string) ([]torrent.Distribution, error) {
	url := p.BaseURL + "/search/" + searchPhrase

	lines := []torrent.Distribution{}

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
		line := torrent.Distribution{}
		firstCol := col.First()
		line.LastUpdated = firstCol.Text()
		secondCol := firstCol.Next()
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

func (p *Parser) ParsePage(id int) torrent.Distribution {
	url := p.BaseURL + "/torrent/" + strconv.Itoa(id)
	line := torrent.Distribution{
		Id: id,
	}

	document, _ := goquery.NewDocument(url)

	content := document.Find("div#content")

	line.MagnetLink, _ = content.
		Find("div#download").
		Find("a").
		First().
		Attr("href")

	details := content.Find("table#details").Find("tr")

	details.Each(func(i int, sel *goquery.Selection) {
		first := sel.Find("td.header")
		switch first.Text() {
		case "Добавлен":
			dateString := first.Next().Text()
			r := regexp.MustCompile("\\d+-\\d+-\\d+ \\d+:\\d+:\\d+")
			line.LastUpdated = r.FindString(dateString)
			break
		case "Размер":
			line.Size = first.Next().Text()
			break

		}
		//println(i, line.Size)
	})

	line.Title = document.Find("div#all").Find("h1").Text()
	line.TopicLink = "/torrent/" + strconv.Itoa(id)
	line.DownloadLink = "/download/" + strconv.Itoa(id)
	//line.LastUpdated = "10"
	return line
}
