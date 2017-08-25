package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"bufio"
)

func getSitesList() [2]string {
	sitesList := [2]string{
		"http://donothingfor2minutes.com/",
		"http://stenadobra.ru/",
		// "http://humandescent.com",
		// "http://thefirstworldwidewebsitewerenothinghappens.com",
		// "http://button.dekel.ru",
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

func createFile(filePath string) (*os.File, string) {
	filePath = "./downloaded/" + filePath
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return file, filePath
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
	file, _ := createFile(fileName)
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

func readFile(filePath string) {
    file, _ := os.Open(filePath)
	f := bufio.NewReader(file)
    for {
        read_line, _ := f.ReadString('\n')
        fmt.Print(read_line)
    }
}

func downloadFile(rawURL string, c chan string) {
	start := time.Now()
	fmt.Println("\n\n")
	fmt.Println(fmt.Sprintf("Start to download %s", rawURL))
	fileName := getFileName(rawURL)
	file, filePath := createFile(fileName)
	defer file.Close()
	resp := getConnection(rawURL)
	defer resp.Body.Close()

	// readFile(resp.Body);
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	end := time.Now()
	fmt.Println(fmt.Sprintf("%s %s with %v bytes downloaded, time=%s", resp.Status, fileName, size, end.Sub(start)))

	c <- filePath
}

func main() {
	fmt.Println("Downloading file...")

	start := time.Now()
	sitesList := getSitesList()
	c := make(chan string, len(sitesList))
	for _, rawURL := range(sitesList) {
		// downloadFileSync(rawURL)
		go downloadFile(rawURL, c)
		// filePath := <- c
		// readFile(filePath);
		fmt.Println("channel length", len(c))
	}

	// close(c)
	for i := range c {
		fmt.Println(fmt.Sprintf("%s done", i))
	}

	end := time.Now()
	fmt.Println('\n')
	fmt.Println(end.Sub(start))
}
