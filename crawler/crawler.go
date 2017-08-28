package main

import (
	"fmt"
	"time"
	"github.com/LepikovStan/ScrapeLinks"
)

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
	chanLength := len(sitesList)
	c := make(chan ScrapeLinks.RefsList, chanLength)

	for _, rawURL := range(sitesList) {
		go ScrapeLinks.Run(rawURL, c)
	}

	result := ScrapeLinks.Format(c, chanLength);
	ScrapeLinks.Print(result);

	end := time.Now()
	fmt.Println("\n")
	fmt.Println(end.Sub(start))
}
