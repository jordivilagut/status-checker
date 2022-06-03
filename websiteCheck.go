package main

import "fmt"

type websiteCheck struct {
	Url        string
	StatusCode int
	Error      error
}

func (w websiteCheck) getStatus() string {
	if w.StatusCode == 200 {
		return "SUCCESS"
	}

	return "UNRESPONSIVE"
}

func (w websiteCheck) print() {
	if w.Error != nil {
		fmt.Println(w.Url, w.getStatus(), w.Error)
	} else {
		fmt.Println(w.Url, w.getStatus())
	}
}
