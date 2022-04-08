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

func (p *Parser) AssertTokenIsHere(position int, expected token.Token) {
	received, err := p.GetToken(position)

	if err != nil {
		panic(err)
	}

	if !expected.Equals(received) {
		panic(NewParserError("expected: " + expected.String() + " received: " + received.String()))
	}
}

// TODO

// func (p *Parser) ParseEqualsExp(position int) (*ParseResult[Exp], error) {
// 	current := parsels
// }

// func (p *Parser) ParseExp(position int) (*ParseResult[Exp], error) {
// 	return p.ParseEqualsExp(position)
// }

func (p *Parser) ParsePrimaryExp(position int) (*ParseResult[Exp], error) {
	tkn, _ := p.GetToken(position)
	if varToken, ok := tkn.(*token.VariableToken); ok {
		name := varToken.Name
		return NewParseResult[Exp](&VariableExp{Variable: Variable{Name: name}}, position+1), nil
	} else if numToken, ok := tkn.(*token.NumberToken); ok {
		val := numToken.Number
		return NewParseResult[Exp](&NumberExp{Number: val}, position+1), nil
	} else { // TODO missing some else ifs abovs
		return nil, NewParserError("Expected primary expression, received: " + tkn.String())
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

func (p *Parser) ParseAdditiveExp(position int) (*ParseResult[Exp], error) {
	current, _ := p.ParsePrimaryExp(position)
	shouldRun := true

	for shouldRun {
		additiveOp, err := p.ParseAdditiveOp(current.Position)
		if err != nil {
			shouldRun = false
		}
		anotherPrimary, err := p.ParsePrimaryExp(additiveOp.Position)
		if err != nil {
			shouldRun = false
		}
		current = NewParseResult[Exp](&OperatorExp{current.Result, additiveOp.Result, anotherPrimary.Result}, anotherPrimary.Position)

	}

	return current, nil
}

func (p *Parser) ParseLessThanExp(position int) (*ParseResult[Exp], error) {
	current, _ := p.ParseAdditiveExp(position)
	shouldRun := true

	for shouldRun {
		p.AssertTokenIsHere(current.Position, &token.LesserToken{})
		other, err := p.ParseAdditiveExp(current.Position + 1)
		if err != nil {
			shouldRun = false
		}
		current = NewParseResult[Exp](&OperatorExp{current.Result, &LessOp{}, other.Result}, other.Position)
	}

	return current, nil
}
