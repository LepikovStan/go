package main

import (
	"fmt"
	"time"
	"strconv"
	"github.com/LepikovStan/ScrapeLinks"
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

func main() {
	fmt.Println("Start...")

	start := time.Now()
	sitesList := getSitesList()
	result := make(map[string][]ScrapeLinks.Ref)
	counter = 1;
	chanLength := len(sitesList)
	c := make(chan ScrapeLinks.RefsList, chanLength)

	for _, rawURL := range(sitesList) {
		go ScrapeLinks.Run(rawURL, c)
	}

	for ref := range(c) {
		result[ref.Url] = ref.Links
		if (chanLength <= 1) {
			close(c)
		}
		chanLength = chanLength - 1
	}

	for url, links := range(result) {
		fmt.Println(fmt.Sprintf("\n%s) %s:\n\n   links:\n", strconv.Itoa(counter), url))
		for _, link := range(links) {
			fmt.Println(fmt.Sprintf("      Href: %s, Title: %s", link.Href, link.Title))
		}
		counter = counter + 1
	}

	end := time.Now()
	fmt.Println('\n')
	fmt.Println(end.Sub(start))
}
