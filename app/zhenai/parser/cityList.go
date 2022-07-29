package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[\da-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//limit := 3
	for _, m := range matches {
		url := m[1]

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(url),
			ParserFunc: ParseCity,
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}

	return result
}
