package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"


)

const (
	nyaaBaseSearch = "http://nyaa.si/?f=0&c=0_0&q="
)

type NyaaClient struct {
	client http.Client
}

func (c NyaaClient) SearchAnime(query string) {
	bodyRes, err := c.client.Get(nyaaBaseSearch + query);
	if err == nil {
		log.Fatal("Error Searching anime")
	}

	body, err := io.ReadAll(bodyRes.Body)
	fmt.Print(body)
}