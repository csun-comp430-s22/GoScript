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

	// fmt.Printf("received: %#v\n", received)

	// fmt.Printf("expected: %#v\n", expected)

	if err != nil {
		panic(err)
	}

	if !expected.Equals(received) {
		panic(NewParserError("expected: " + expected.String() + " received: " + received.String()))
	}
}

func (p *Parser) ParsePrimaryExp(position int) (*ParseResult[Exp], error) {
	tkn, _ := p.GetToken(position)
	if varToken, ok := tkn.(*token.VariableToken); ok {
		name := varToken.Name
		return NewParseResult[Exp](&VariableExp{Variable: Variable{Name: name}}, position+1), nil
	} else if intToken, ok := tkn.(*token.IntLiteralToken); ok {
		val := intToken.Number
		return NewParseResult[Exp](&IntLiteralExp{Number: val}, position+1), nil
	} else { // TODO missing some else ifs abovs
		return nil, NewParserError("Expected primary expression, received: " + tkn.String())
	}
}

func (p *Parser) ParseAdditiveOp(position int) (*ParseResult[Operator], error) {
	// TODO handle error
	tkn, _ := p.GetToken(position)

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
	// fmt.Printf("position: %v\n", position)

	tkn, err := p.GetToken(position)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("tkn: %v\n", tkn)

	// position will be moved inside the parser function to next or next after usually
	if _, ok := tkn.(*token.IfToken); ok {
		return p.parseSelectionStmt(position)
	} else if _, ok := tkn.(*token.LeftCurlyToken); ok {
		return p.parseBlockStmt(position)
	} else if _, ok := tkn.(*token.FnToken); ok {
		res, err := p.ParseFunctionDefinition(position)
		// this was done cause p.ParseFunctionDefinition(position) returns *ParseResult[*FunctionDef]
		// instead of *ParseResult[Stmt]
		return NewParseResult[Stmt](res.Result, res.Position), err
	} else {
		return nil, NewParserError("expected statement, received: " + tkn.String())
	}
}

// if (exp) stmt | if (exp) stmt else stmt
func (p *Parser) parseSelectionStmt(position int) (*ParseResult[Stmt], error) {
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

func (p *Parser) parseBlockStmt(position int) (*ParseResult[Stmt], error) {
	smts := []Stmt{}
	currentPosition := position + 1
	shouldRun := true
	for shouldRun {
		// tkn, _ := p.GetToken(position)
		// if _, ok := tkn.(*token.RightCurlyToken); ok {
		// 	currentPosition += 1
		// }

		stmt, err := p.ParseStmt(currentPosition)
		if err != nil {
			shouldRun = false
			break
		}
		smts = append(smts, stmt.Result)
		currentPosition = stmt.Position
	}
	return NewParseResult[Stmt](NewBlockStmt(smts), currentPosition), nil
}

func (p *Parser) parseType(position int) (*ParseResult[Type], error) {
	tkn, err := p.GetToken(position)

	if err != nil {
		fmt.Println(err)
	}

	if _, ok := tkn.(*token.IntToken); ok {
		return NewParseResult[Type](&IntType{}, position+1), nil
	} else {
		return nil, NewParserError("super descritptive error")
	}

}

func (p *Parser) parseFuncArgs(position int) ([]*Vardec, int, error) {
	shouldRun := true
	varDecs := []*Vardec{}
	currentPos := position
	for shouldRun {
		tkn, err := p.GetToken(currentPos)
		if err != nil {
			shouldRun = false
			break
		}

		if varTkn, ok := tkn.(*token.VariableToken); ok {

			varType, err := p.parseType(currentPos + 1)
			if err != nil {
				return nil, currentPos, NewParserError(err.Error())
			}
			varDec := NewVardec(&Variable{Name: varTkn.Name}, varType.Result)
			varDecs = append(varDecs, varDec)
			currentPos += 2

			tkn, err := p.GetToken(currentPos)
			// fmt.Printf("tkn: %#v\n", tkn)

			if _, ok := tkn.(*token.CommaToken); err != nil || !ok {
				shouldRun = false
				break
			}
			currentPos += 1

		} else {
			shouldRun = false
			break
		}

	}
	return varDecs, currentPos, nil
}

func (p *Parser) ParseFunctionDefinition(position int) (*ParseResult[*FunctionDef], error) {
	// fmt.Println("FUNCTION")

	returnTypeRes, err := p.parseType(position + 1)
	if err != nil {
		return nil, NewParserError("expected type here")
	}

	funcNameTkn, _ := p.GetToken(returnTypeRes.Position)

	funcName := ""
	if castFuncNameTkn, ok := funcNameTkn.(*token.VariableToken); ok {
		funcName = castFuncNameTkn.Name
	}

	p.AssertTokenIsHere(returnTypeRes.Position+1, &token.LeftParenToken{})

	args, position, _ := p.parseFuncArgs(returnTypeRes.Position + 2)

	if err != nil {
		return nil, NewParserError("oof")
	}

	p.AssertTokenIsHere(position, &token.RightParenToken{})

	stmtRes, err := p.ParseStmt(position + 1)

	if err != nil {
		return nil, err
	}

	return NewParseResult(NewFunctionDef(NewFunctionName(funcName), args, stmtRes.Result, returnTypeRes.Result), stmtRes.Position), nil
}

// // Old
// func (p *Parser) parseProgram(position int) (*ParseResult[*Program], error) {
// 	stmt, _ := p.ParseStmt(position)

// 	return NewParseResult(NewProgram(stmt.Result), stmt.Position), nil
// }

func (p *Parser) parseProgram(position int) (*ParseResult[*Program], error) {

	shouldRun := true

	funcDefs := []*FunctionDef{}
	currentPos := position
	for shouldRun {
		funcDef, err := p.ParseFunctionDefinition(currentPos)
		if err != nil {
			shouldRun = false
			break
		}
		funcDefs = append(funcDefs, funcDef.Result)
		currentPos = funcDef.Position

		if currentPos == len(p.Tokens)-1 {
			// if reached the end leave currentPos as is
			shouldRun = false
			break

		} else {
			// no errors and next token is fn token
			p.AssertTokenIsHere(currentPos+1, &token.FnToken{})
			currentPos += 1
		}
	}

	return NewParseResult(NewProgram(funcDefs), currentPos), nil
}

func (p *Parser) ParseProgram() (*Program, error) {
	program, _ := p.parseProgram(0)

	if program.Position == len(p.Tokens)-1 {
		return program.Result, nil
	} else {
		fmt.Printf("program.Position: %v\n", program.Position)
		panic("Remaining tokens at end")
	}
}

func (p *Parser) parseFuncParams(position int) ([]Exp, int, error) {
	shouldRun := true
	params := []Exp{}
	currentPos := position
	for shouldRun {
		exp, err := p.ParsePrimaryExp(currentPos)
		if err != nil {
			shouldRun = false
			break
		}

		params = append(params, exp.Result)
		currentPos = exp.Position

		tkn, err := p.GetToken(currentPos)

		if _, ok := tkn.(*token.CommaToken); err != nil || !ok {
			shouldRun = false
			break
		}
		currentPos += 1

	}
	return params, currentPos, nil
}

func (p *Parser) ParseFunctionCall(position int) (*ParseResult[*FunctionCallExp], error) {
	funcNameTkn, _ := p.GetToken(position)
	funcName := ""
	if castFuncNameTkn, ok := funcNameTkn.(*token.VariableToken); ok {
		funcName = castFuncNameTkn.Name
	}

	p.AssertTokenIsHere(position+1, &token.LeftParenToken{})

	params, position, _ := p.parseFuncParams(position + 2)

	p.AssertTokenIsHere(position, &token.RightParenToken{})

	return NewParseResult(NewFunctionCallExp(NewFunctionName(funcName), params), position), nil
}

/* func (p *Parser) ParseHigherOrderFunctionDef (position int) (*ParseResult[*HigherOrderFunctionDef], error) {
	returnTypeRes, err := p.parseType(position + 1)
	if err != nil {
		return nil, NewParserError("expected type here")
	}

	funcNameTkn, _ := p.GetToken(returnTypeRes.Position)

	funcName := ""
	if castFuncNameTkn, ok := funcNameTkn.(*token.VariableToken); ok {
		funcName = castFuncNameTkn.Name
	}

	p.AssertTokenIsHere(returnTypeRes.Position+1, &token.LeftParenToken{})

	params, position, _ := p.parseFuncParams(returnTypeRes.Position + 2)

	if err != nil {
		return nil, NewParserError("oof")
	}

	p.AssertTokenIsHere(position, &token.RightParenToken{})

	stmtRes, err := p.ParseStmt(position + 1)

	if err != nil {
		return nil, err
	}

	return NewParseResult(NewHigherOrderFunctionDef(NewFunctionName(funcName), params, stmtRes.Result, returnTypeRes.Result), stmtRes.Position), nil
}
*/

/* func (p *Parser) ParseHigherOrderFunctionCall (position int) (*ParseResult[*HigherOrderFunctionCallExp], error) {
	funcNameTkn, _ := p.GetToken(position)
	funcName := ""
	if castFuncNameTkn, ok := funcNameTkn.(*token.VariableToken); ok {
		funcName = castFuncNameTkn.Name
	}

	p.AssertTokenIsHere(position+1, &token.LeftParenToken{})

	params, position, _ := p.parseFuncParams(position + 2)

	p.AssertTokenIsHere(position, &token.RightParenToken{})

	return NewParseResult(NewHigherOrderFunctionCallExp(NewFunctionName(funcName), params), position), nil
}
*/
