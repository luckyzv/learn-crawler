package engine

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParserFunc func([]byte, string) ParseResult

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
