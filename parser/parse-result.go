package parser

// go doesn't really have generic programming yet
// so we can use empty interface to take in anything

type ParseResult struct {
	Result   interface{}
	Position int
}

func NewParseResult(result interface{}, position int) *ParseResult {
	return &ParseResult{Result: result, Position: position}
}
