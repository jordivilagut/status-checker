package main

import "net/http"

func statusCode(response *http.Response) int {
	if response == nil {
		return 0
	}
	return response.StatusCode
}
