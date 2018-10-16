package engine

type ParserFunc func(contents []byte,url string) ParserResult

type Parser interface {
	Parse(contents []byte,url string) ParserResult
	Serialize() (name string,args interface{})
} 
type Request struct {
	Url string
	Parser Parser
}
type ParserResult struct {
	Requests []Request
	//Items []interface{}
	Items []Item

}

type Item struct {

	Url string
	Id string
	Type string
	Payload interface{}
}
type NilParser struct {

}

func (n NilParser) Parse(_ []byte, url string) ParserResult {
	return ParserResult{}
}

func (n NilParser) Serialize() (name string,args interface{}) {
	return "NilParser",nil
}

type FuncParser struct {
	parser ParserFunc
	Name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.parser(contents,url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name,nil
}

func NewFuncParser(p ParserFunc,name string)  *FuncParser  {

	return &FuncParser{
		parser:p,
		Name:name,
	}
}