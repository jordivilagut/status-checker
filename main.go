package main

import (
	"net/http"
	"time"
)

var sites = []string{
	"https://google.com",
	"https://facebook.com",
	"https://stackoverflow.com",
	"https://golang.com",
	"https://cjlsxnspd.com/",
	"https://amazon.com",
	"https://jenkins.andjoy.life/",
}

func main() {

	c := make(chan websiteCheck)

	for _, site := range sites {
		go checkStatus(site, c)
	}

	for wc := range c {
		go func(check websiteCheck) {
			time.Sleep(2 * time.Second)
			checkStatus(check.Url, c)
		}(wc)
	}
}

func checkStatus(site string, c chan websiteCheck) {
	response, err := http.Get(site)
	check := websiteCheck{Url: site, StatusCode: statusCode(response), Error: err}
	check.print()
	c <- check
}
