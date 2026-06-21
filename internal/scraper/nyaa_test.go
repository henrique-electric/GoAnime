package scraper

import (
	"testing"
)

func TestNyaaSearch(t *testing.T) {
	var c NyaaClient
	c.SearchAnime("High school")
}