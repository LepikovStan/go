package main

import (
	"fmt"
	"strings"
	"time"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)

var counter int;

func getSitesList() [21]string {
	sitesList := [21]string{
		"http://donothingfor2minutes.com/",
		"http://stenadobra.ru/",
		"http://humandescent.com",
		"http://thefirstworldwidewebsitewerenothinghappens.com",
		"http://button.dekel.ru",
		"http://www.randominio.com/",
		"http://thenicestplaceontheinter.net/",
		"http://www.catsthatlooklikehitler.com/",
		"http://www.thefirstworldwidewebsitewerenothinghappens.com/",
		"http://www.donothingfor2minutes.com/",
		"http://www.howmanypeopleareinspacerightnow.com/",
		"http://www.humanclock.com/",
		"http://fucking-great-advice.ru/",
		"http://www.cesmes.fi/pallo.swf",
		"http://button.dekel.ru/",
		"http://www.rainfor.me/",
		"http://loudportraits.com/",
		"http://sprosimamu.ru/",
		"http://www.bandofbridges.com/",
		"http://www.catsboobs.com/",
		"http://www.incredibox.com/",
	}
	return sitesList
}

type Refs struct {
	title string
	href string
	url string
}

type R struct {
	url string
	links []Refs
}

func ScrapeLinks(url string, c chan R) {
	var links []Refs

 	doc, err := goquery.NewDocument(url)

 	if err != nil {
 		panic(err)
 	}

 	doc.Find("a").Each(func(i int, s *goquery.Selection) {
 		Title := strings.TrimSpace(s.Text())
		Href, _ := s.Attr("href")
		ref := Refs{
			title: Title,
			href: Href,
			url: url,
	    }
		links = append(links, ref)
 	})
	res := R{
		url: url,
		links: links,
	}

	c <- res
 }

func main() {
	fmt.Println("Start...")

	start := time.Now()
	sitesList := getSitesList()
	result := make(map[string][]Refs)
	counter = 1;
	chanLength := len(sitesList)
	c := make(chan R, chanLength)

	for _, rawURL := range(sitesList) {
		go ScrapeLinks(rawURL, c)
	}

	for ref := range(c) {
		result[ref.url] = ref.links
		if (chanLength <= 1) {
			close(c)
		}
		chanLength = chanLength - 1
	}

	for url, links := range(result) {
		fmt.Println(fmt.Sprintf("\n%s) %s:\n\n   links:\n", strconv.Itoa(counter), url))
		for _, link := range(links) {
			fmt.Println(fmt.Sprintf("      Href: %s, Title: %s", link.href, link.title))
		}
		counter = counter + 1
	}

	end := time.Now()
	fmt.Println('\n')
	fmt.Println(end.Sub(start))
}
