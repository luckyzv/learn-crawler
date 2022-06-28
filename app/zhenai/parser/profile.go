package parser

import (
	"crawler/app/model"
	"crawler/engine"
	"regexp"
)

var importantProfileRe = regexp.MustCompile(`class="m-btn purple"[^>]*>([^<]+)</div>`)
var extraProfileRe = regexp.MustCompile(`class="m-btn pink"[^>]*>([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{NickName: name}
	formatImportantProfile(contents, importantProfileRe, &profile)
	formatExtraProfile(contents, extraProfileRe, &profile)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func formatImportantProfile(contents []byte, re *regexp.Regexp, profile *model.Profile) {
	match := re.FindAllSubmatch(contents, -1)

	profile.Marriage = string(match[0][1])
	profile.Age = string(match[1][1])
	profile.Star = string(match[2][1])
	profile.Height = string(match[3][1])
	profile.Weight = string(match[4][1])
	profile.WorkPlace = string(match[5][1])
	profile.Income = string(match[6][1])
	profile.Job = string(match[7][1])
	profile.Education = string(match[8][1])

}

func formatExtraProfile(contents []byte, re *regexp.Regexp, profile *model.Profile) {
	match := re.FindAllSubmatch(contents, -1)
	profile.HouKou = string(match[1][1])
}
