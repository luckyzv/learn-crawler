package engine

import (
	"crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)

	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch error: fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(contents, r.Url), nil
}
