package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)

		contents, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch error: fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(contents)
		requests = append(requests, parseResult.Requests...)

		for _, v := range parseResult.Items {
			fmt.Println("items: ", v)
		}
	}

}
