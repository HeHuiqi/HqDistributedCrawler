package parser

import (
	"HqDistributedCrawler/engine"
	"regexp"
	"strconv"
	"HqDistributedCrawler/model"
)

const profileRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var ageRe  = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var sexRe  = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var HeightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe  = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var carRe  = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var houseRe  = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)

var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe  = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
func ParserProfile(contents []byte,name string,url string) engine.ParserResult  {


	profile := model.Profile{}

	profile.Name = name
	profile.Gender = extractString(contents,sexRe)

	age,err:= strconv.Atoi(extractString(contents,ageRe))
	if err != nil {
		age = 0
	}
	profile.Age = age
	height,err := strconv.Atoi(extractString(contents,HeightRe))
	if err != nil {
		height = 0
	}
	profile.Height = height

	profile.Marriage = extractString(contents,marriageRe)

	profile.Income = extractString(contents,incomeRe)

	profile.Car = extractString(contents,carRe)
	profile.House = extractString(contents,houseRe)

	profile.Hukou = extractString(contents,hukouRe)

	profile.Xinzuo = extractString(contents,xinzuoRe)

	profile.Occupation = extractString(contents,occupationRe)

	result := engine.ParserResult{
		Items:[]engine.Item{
			{
				Url:url,
				Type:"zhenai",
				Id:extractString([]byte(url),idUrlRe),
				Payload:profile,
			},
		},
	}


	/*
	var guessRe = regexp.MustCompile(``)
	matchs := guessRe.FindAllSubmatchIndex(contents,-1)
	for _,m := range matchs {
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(result.Requests,engine.Request{
			Url:url,
			Parser: NewProfileParser(config.ParserProfile),
		})
	}
	*/



	return result
}
func extractString(contents []byte,regex *regexp.Regexp) string  {

	match := regex.FindSubmatch(contents)

	if len(match) >=2  {
		return string(match[1])

	}
	return ""
}

type ProfileParser struct {
	UserName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParserResult {

	return ParserProfile(contents,p.UserName,url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParserProfile",p.UserName
}

func NewProfileParser(name string) *ProfileParser  {
	return &ProfileParser{UserName:name}
}

//func ProfileParser(name string) engine.ParserFunc  {
//	return func(c []byte,url string) engine.ParserResult {
//		return ParserProfile(c,name,url)
//	}
//}