package parser

import (
	"fmt"
	"regexp"
	"yq.com/go_study/crawler/engine"
)

const citylistRe = `<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(citylistRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, m[1] )
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[0]),
			ParserFunc: nil,
		})
		fmt.Printf("City: %s, URL: %s\n", m[1], m[0])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
