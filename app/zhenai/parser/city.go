package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := m[1]
		name := m[2]

		result.Items = append(result.Items, "User "+string(name))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(url),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(name))
			},
		})
	}

	return result
}
