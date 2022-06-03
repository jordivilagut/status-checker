package main

import (
	"net/http"
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

	for i := 0; i < len(sites); i++ {
		check := <-c
		check.print()
	}
}

func checkStatus(site string, c chan websiteCheck) {
	response, err := http.Get(site)
	c <- websiteCheck{Url: site, StatusCode: statusCode(response), Error: err}
}
