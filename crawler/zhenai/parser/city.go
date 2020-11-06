package parser

import (
	"learn-go/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(context []byte) engine.ParseResult {
				return ParseProfile(contents, string(m[2]))
			},
		})
	}
	return result
}