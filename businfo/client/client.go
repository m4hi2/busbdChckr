package client

import "net/http"

var client *http.Client

func GetClient() *http.Client {
	if client != nil {
		return client
	}
	return CreateClient()
}

func CreateClient() *http.Client {
	client = http.DefaultClient
	return client
}
