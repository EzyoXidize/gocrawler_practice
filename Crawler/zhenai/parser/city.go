package parser

import (
	"practice/Crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	// 用户信息正则
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	// 城市连接正则 , 包括翻页
	)
func ParseCity(contents []byte) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _,m := range matches {
		name := string(m[2])

		result.Requests = append(result.Requests, engine.Request{
				Url			:	string(m[1]),
				ParserFunc	: func(c []byte) engine.ParseResult {
					return  ParserProfile(c,name)
				},
			})
	}

	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _, m := range matches {
		result.Requests = append(result.Requests,engine.Request{
			Url			: 	string(m[1]),
			ParserFunc	:	ParseCity,
		})
	}
	return result

}