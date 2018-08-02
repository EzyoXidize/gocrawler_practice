package parser

import (
	"practice/Crawler/engine"
	"regexp"
	"strconv"
	"practice/Crawler/model"
)

var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var ageRe  = regexp.MustCompile(` <td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe  =  regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)


func ParserProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name


	age, err := strconv.Atoi(extractString(contents,ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Weight = extractString(contents,weightRe)
	profile.Marriage = extractString(contents,marriageRe)
	profile.Gender = extractString(contents,genderRe)
	profile.Height = extractString(contents,heightRe)
	profile.Income = extractString(contents,incomeRe)
	profile.Occupation = extractString(contents,occupationRe)
	profile.Xinzuo= extractString(contents,xingzuoRe)
	profile.Education = extractString(contents,educationRe)
	profile.Hokou = extractString(contents,hokouRe)


	result := engine.ParseResult{
		Items:[]interface{} {profile} ,
	}

	return result
}

func extractString(contents []byte , re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}else {
		return ""
	}
}

