package parser

import (
	"fmt"
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
		posStr := strconv.Itoa(position)
		return nil, NewParserError("invalid token position: " + posStr)
	}
}

func (p *Parser) AssertTokenIsHere(position int, expected token.Token) {
	received, err := p.GetToken(position)

	fmt.Printf("received: %#v\n", received)

	fmt.Printf("expected: %#v\n", expected)

	if err != nil {
		panic(err)
	}

	if !expected.Equals(received) {
		panic(NewParserError("expected: " + expected.String() + " received: " + received.String()))
	}
}

func (p *Parser) ParseEqualsExp(position int) (*ParseResult[Exp], error) {
	current, _ := p.ParseComparisonExp(position)
	shouldRun := true

	for shouldRun {
		p.AssertTokenIsHere(current.Position, &token.EqualsToken{})
		other, err := p.ParseComparisonExp(current.Position + 1)
		if err != nil {
			shouldRun = false
			break
		}
		current = NewParseResult[Exp](NewOpExp(current.Result, &EqualsOp{}, other.Result), other.Position)

	}

	return current, nil
}

func (p *Parser) ParseExp(position int) (*ParseResult[Exp], error) {
	return p.ParseEqualsExp(position)
}

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
	} else if _, ok := tkn.(*token.ModToken); ok {
		return NewParseResult[Operator](&ModOp{}, position+1), nil
	} else if _, ok := tkn.(*token.PowerToken); ok {
		return NewParseResult[Operator](&PowOp{}, position+1), nil
	} else if _, ok := tkn.(*token.PipeOperatorToken); ok {
		return NewParseResult[Operator](&PipeOp{}, position+1), nil
	} else if _, ok := tkn.(*token.OrToken); ok {
		return NewParseResult[Operator](&OrOp{}, position+1), nil
	} else if _, ok := tkn.(*token.AndToken); ok {
		return NewParseResult[Operator](&AndOp{}, position+1), nil
	} else if _, ok := tkn.(*token.AssignmentToken); ok {
		return NewParseResult[Operator](&AssignmentOp{}, position+1), nil
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
			break
		}
		anotherPrimary, err := p.ParsePrimaryExp(additiveOp.Position)
		if err != nil {
			shouldRun = false
			break
		}
		current = NewParseResult[Exp](NewOpExp(current.Result, additiveOp.Result, anotherPrimary.Result), anotherPrimary.Position)

	}

	return current, nil
}

func (p *Parser) ParseComparisonOp(position int) (*ParseResult[Operator], error) {
	tkn, _ := p.GetToken(position)

	if _, ok := tkn.(*token.EqualsToken); ok {
		return NewParseResult[Operator](&EqualsOp{}, position+1), nil
	} else if _, ok := tkn.(*token.NotEqualsToken); ok {
		return NewParseResult[Operator](&NotEqualsOp{}, position+1), nil
	} else if _, ok := tkn.(*token.LesserToken); ok {
		return NewParseResult[Operator](&LessOp{}, position+1), nil
	} else if _, ok := tkn.(*token.GreaterToken); ok {
		return NewParseResult[Operator](&GreaterOp{}, position+1), nil
	} else {
		return nil, NewParserError("super descritptive error")
	}

}

func (p *Parser) ParseComparisonExp(position int) (*ParseResult[Exp], error) {
	additive, err := p.ParseAdditiveExp(position)
	if err != nil {
		return nil, NewParserError(err.Error())
	}
	comparisonOp, err := p.ParseComparisonOp(additive.Position)
	if err != nil {
		return nil, NewParserError(err.Error())
	}
	secAdditive, err := p.ParseAdditiveExp(comparisonOp.Position)
	if err != nil {
		return nil, NewParserError(err.Error())
	}
	return NewParseResult[Exp](NewOpExp(additive.Result, comparisonOp.Result, secAdditive.Result), secAdditive.Position), nil
}

func (p *Parser) ParseLogicalOp(position int) (*ParseResult[Operator], error) {
	tkn, err := p.GetToken(position)

	fmt.Println(err)

	if _, ok := tkn.(*token.AndToken); ok {
		return NewParseResult[Operator](&AndOp{}, position+1), nil
	} else if _, ok := tkn.(*token.OrToken); ok {
		return NewParseResult[Operator](&OrOp{}, position+1), nil
	} else if _, ok := tkn.(*token.NegateToken); ok {
		return NewParseResult[Operator](&NegateOp{}, position+1), nil
	} else {
		return nil, NewParserError("super descritptive error")
	}
}

func (p *Parser) ParseLogicalExp(position int) (*ParseResult[Exp], error) {
	current, _ := p.ParsePrimaryExp(position)
	shouldRun := true

	for shouldRun {
		logicalOp, err := p.ParseAdditiveOp(current.Position)
		if err != nil {
			shouldRun = false
			break
		}
		anotherPrimary, err := p.ParsePrimaryExp(logicalOp.Position)
		if err != nil {
			shouldRun = false
			break
		}
		current = NewParseResult[Exp](NewOpExp(current.Result, logicalOp.Result, anotherPrimary.Result), anotherPrimary.Position)

	}

	return current, nil
}

func (p *Parser) ParseStmt(position int) (*ParseResult[Stmt], error) {
	fmt.Printf("position: %v\n", position)

	tkn, err := p.GetToken(position)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("tkn: %v\n", tkn)

	if _, ok := tkn.(*token.IfToken); ok {
		return p.parseSelectionStatement(position)

	} else if _, ok := tkn.(*token.LeftCurlyToken); ok {
		smts := []Stmt{}
		currentPosition := position + 1
		shouldRun := true
		for shouldRun {
			stmt, err := p.ParseStmt(currentPosition)
			if err != nil {
				shouldRun = false
				break
			}
			smts = append(smts, stmt.Result)
			currentPosition = stmt.Position
		}

		return NewParseResult[Stmt](NewBlockStmt(smts), currentPosition), nil

	} else {
		return nil, NewParserError("expected statement, received: " + tkn.String())
	}
}

// if (exp) stmt | if (exp) stmt else stmt
func (p *Parser) parseSelectionStatement(position int) (*ParseResult[Stmt], error) {
	p.AssertTokenIsHere(position+1, &token.LeftParenToken{})
	guard, _ := p.ParseComparisonExp(position + 2)
	p.AssertTokenIsHere(guard.Position, &token.RightParenToken{})
	trueBranch, _ := p.ParseStmt(guard.Position + 1)
	tkn, _ := p.GetToken(trueBranch.Position + 1)
	if _, ok := (tkn).(*token.ElseToken); ok {
		falseBranch, _ := p.ParseStmt(trueBranch.Position + 2)
		return NewParseResult[Stmt](NewIfElseStmt(guard.Result, trueBranch.Result, falseBranch.Result), falseBranch.Position), nil
	}

	return NewParseResult[Stmt](NewIfStmt(guard.Result, trueBranch.Result), trueBranch.Position), nil
}
