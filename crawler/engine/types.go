package engine

import "learn-go/crawler_distributed/config"

type ParseFunc func(
	contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct {
}

func (n NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (
	name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
