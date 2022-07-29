package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])

		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: profileParser(name),
		})
	}

	return result
}

func profileParser(name string) engine.ParserFunc {
	return func(bytes []byte, url string) engine.ParseResult {
		return ParseProfile(bytes, name, url)
	}
}
