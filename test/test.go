package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func getSitesList() [5]string {
	sitesList := [5]string{
		"http://donothingfor2minutes.com/",
		"http://stenadobra.ru/",
		"http://humandescent.com",
		"http://thefirstworldwidewebsitewerenothinghappens.com",
		"http://button.dekel.ru",
	}
	return sitesList
}

func getFileName(rawURL string) string {
	fileURL, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	host := fileURL.Host
	segments := strings.Split(host, ".")
	domain := segments[0]
	fileName := domain + ".html"
	return fileName
}

func createFile(filePath string) *os.File {
	filePath = "./downloaded/" + filePath
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return file
}

func getConnection(rawURL string) *http.Response {
	connect := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := connect.Get(rawURL) // add a filter to check redirect

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return resp
}

func downloadFileSync(rawURL string) {
	start := time.Now()
	fmt.Println("\n\n")
	fmt.Println(fmt.Sprintf("Start to download %s", rawURL))
	fileName := getFileName(rawURL)
	file := createFile(fileName)
	defer file.Close()
	resp := getConnection(rawURL)

	fmt.Println(resp.Status)
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	end := time.Now()
	fmt.Printf("%s with %v bytes downloaded, time=%s", fileName, size, end.Sub(start))
}

func downloadFile(rawURL string, c chan string) {
	start := time.Now()
	fmt.Println("\n\n")
	fmt.Println(fmt.Sprintf("Start to download %s", rawURL))
	fileName := getFileName(rawURL)
	file := createFile(fileName)
	defer file.Close()
	resp := getConnection(rawURL)

	fmt.Println(resp.Status)
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	end := time.Now()
	result := fmt.Sprintf("%s with %v bytes downloaded, time=%s", fileName, size, end.Sub(start))

	c <- result
}

func main() {
	fmt.Println("Downloading file...")

	start := time.Now()
	c := make(chan string)
	sitesList := getSitesList()
	for _, rawURL := range(sitesList) {
		// downloadFileSync(rawURL)
		go downloadFile(rawURL, c)
		fmt.Println(fmt.Sprintf("\n %s", <- c))
	}

	end := time.Now()
	fmt.Println('\n')
	fmt.Println(end.Sub(start))
}
