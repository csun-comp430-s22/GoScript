package parser

import (
	"strconv"

	"github.com/vSterlin/goscript/token"
)

type Parser struct {
	Tokens []token.Token
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{Tokens: tokens}
}

func (p *Parser) GetToken(position int) (token.Token, error) {

	// to be able to return uninitialized token as nil
	if position >= 0 && position < len(p.Tokens) {
		return p.Tokens[position], nil
	} else {
		var t token.Token
		posStr := strconv.Itoa(position)
		return t, NewParserError("invalid token position: " + posStr)
	}
}

func (p *Parser) ParseAdditiveOp(position int) (*ParseResult[Operator], error) {
	// TODO handle error
	tkn, _ := p.GetToken(position)
	// tokenType = reflect.TypeOf(tkn)

	// cast tkn as PlusToken, if casted then it's instance of it
	if _, ok := tkn.(*token.PlusToken); ok {
		return NewParseResult[Operator](&PlusOp{}, position+1), nil
	} else if _, ok := tkn.(*token.MinusToken); ok {
		return NewParseResult[Operator](&MinusOp{}, position+1), nil
	} else if _, ok := tkn.(*token.MultToken); ok {
		return NewParseResult[Operator](&MultiplyOp{}, position+1), nil
	} else if _, ok := tkn.(*token.DivToken); ok {
		return NewParseResult[Operator](&DivideOp{}, position+1), nil
	} else if _, ok := tkn.(*token.PipeOperatorToken); ok {
		return NewParseResult[Operator](&PipeOp{}, position+1), nil
	} else {
		return nil, NewParserError("err")
	}

}
