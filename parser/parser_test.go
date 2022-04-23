package parser

import (
	"fmt"

	"testing"

	"github.com/vSterlin/goscript/token"
)

// helper to take string and return tokens
func tokenize(input string) []token.Token {
	tokenizer := token.NewTokenizer(input)
	return tokenizer.Tokenize()
}

func TestPlusOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &PlusOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &PlusOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &PlusOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestAdditiveOpPlus(t *testing.T) {

	tokens := []token.Token{&token.PlusToken{}, &token.IntLiteralToken{Number: 1}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&PlusOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}
func TestMinusOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &MinusOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &MinusOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &MinusOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestSubtractiveOpMinus(t *testing.T) {

	tokens := []token.Token{&token.MinusToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&MinusOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}
func TestDivideOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &DivideOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &DivideOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &DivideOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestDivisionOpDivide(t *testing.T) {

	tokens := []token.Token{&token.DivToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&DivideOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

func TestMultiplyOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &MultiplyOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &MultiplyOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &MultiplyOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestMultiplicationOpMultiply(t *testing.T) {

	tokens := []token.Token{&token.MultToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&MultiplyOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

func TestPowOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &PowOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &PowOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &PowOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestExponentOpPower(t *testing.T) {

	tokens := []token.Token{&token.PowerToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&PowOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

func TestModOpTestExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &ModOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &ModOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &ModOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestModuloOpMod(t *testing.T) {

	tokens := []token.Token{&token.ModToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&ModOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

func TestAndOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &AndOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &AndOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &AndOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}
}

/* func TestLessThanExp(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{1}, &token.LesserToken{}, &token.IntLiteralToken{2}}
	parser := NewParser(tokens)

	parseResult, err := parser.ParseComparisonExp(0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	expected := NewParseResult(NewOpExp(&IntLiteralExp{1}, &LessOp{}, &IntLiteralExp{2}), 3)

	fmt.Printf("parseResult: %v\n", parseResult)
	fmt.Printf("expected: %v\n", expected)

	if !parseResult.Equals(expected) {
		t.Error("Expected parse result did not equal actual")

	}
} */

func TestPipeOperatorOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &PipeOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &PipeOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &PipeOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestPipeOperatorOpPipeOp(t *testing.T) {

	tokens := []token.Token{&token.PipeOperatorToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&PipeOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}
func TestEqualsOp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &EqualsOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &EqualsOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &EqualsOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}
func TestNotEquals(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &NotEqualsOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &NotEqualsOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &NotEqualsOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestOrOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &OrOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &OrOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &OrOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}

}

func TestOrOpMod(t *testing.T) {
	tokens := []token.Token{&token.OrToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&OrOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

func TestLessThanOp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &LessOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &LessOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &LessOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}
}

func TestAndOp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &AndOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &AndOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &AndOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}
}

func TestAmpersandOpAnd(t *testing.T) {
	tokens := []token.Token{&token.AndToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseLogicalOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&AndOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

func TestAssignmentOp(t *testing.T) {
	tokens := []token.Token{&token.AssignmentToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&AssignmentOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

/*func TestBlockStmtOp(t *testing.T) {
	tokens := []token.Token{&token.BlockToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseBlockStmt(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&BlockStmt{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}
}*/
func TestGreaterThanOp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &GreaterOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &GreaterOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &GreaterOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}
}
func TestLessEqualExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &LessEqualOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &LessEqualOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &LessEqualOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}
}

//TODO: Merged Tokens needed for this parser
/* func TestGreaterEqualOpExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &GreaterEqualOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &GreaterEqualOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &GreaterEqualOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}
}
func TestGreaterEqualModTest(t *testing.T) {
	tokens := []token.Token{&token.GreaterToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseComparisonOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}
	if !parseResult.Equals(NewParseResult[Operator](&GreaterEqualOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")
	}
} */

func TestNegateOpTestExp(t *testing.T) {
	first := NewOpExp(&IntLiteralExp{1}, &NegateOp{}, &IntLiteralExp{1})
	second := NewOpExp(&IntLiteralExp{1}, &NegateOp{}, &IntLiteralExp{1})
	third := NewOpExp(&IntLiteralExp{1}, &NegateOp{}, &IntLiteralExp{2})

	fmt.Printf("%#v, %#v", first, second)
	if !first.Equals(second) {
		t.Error("Expected first and second operator expression to be equal")
	}
	if !first.Equals(first) {
		t.Error("Expected first operator expression to be equal to itself")
	}
	if first.Equals(third) {
		t.Error("Expected first operator expression not to be equal to third")
	}
}

func TestNegateOpMod(t *testing.T) {
	tokens := []token.Token{&token.NegateToken{}, &token.ModToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseLogicalOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}
	if !parseResult.Equals(NewParseResult[Operator](&NegateOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")
	}
}

func TestIncorrectPosition(t *testing.T) {
	parser := NewParser([]token.Token{})
	_, err := parser.GetToken(-1)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestAdditivePlusExp(t *testing.T) {
	tokens := []token.Token{&token.VariableToken{Name: "x"}, &token.EqualsToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseAdditiveExp(0)

	if parseResult.Equals(NewParseResult[Exp](NewOpExp(&VariableExp{}, &EqualsOp{}, &IntLiteralExp{1}), 3)) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestAdditiveExp(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.PlusToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseAdditiveExp(0)

	if !parseResult.Equals(NewParseResult[Exp](NewOpExp(&IntLiteralExp{1}, &PlusOp{}, &IntLiteralExp{1}), 3)) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestComparisonExpLess(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.LesserToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseComparisonExp(0)

	if !parseResult.Equals(NewParseResult[Exp](NewOpExp(&IntLiteralExp{1}, &LessOp{}, &IntLiteralExp{1}), 3)) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestComparisonExpGreater(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.GreaterToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseComparisonExp(0)

	if !parseResult.Equals(NewParseResult[Exp](NewOpExp(&IntLiteralExp{1}, &GreaterOp{}, &IntLiteralExp{1}), 3)) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestComparisonExpEquals(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.EqualsToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseComparisonExp(0)

	if !parseResult.Equals(NewParseResult[Exp](NewOpExp(&IntLiteralExp{1}, &EqualsOp{}, &IntLiteralExp{1}), 3)) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestComparisonExpNotEquals(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.NotEqualsToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseComparisonExp(0)

	if !parseResult.Equals(NewParseResult[Exp](NewOpExp(&IntLiteralExp{1}, &NotEqualsOp{}, &IntLiteralExp{1}), 3)) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestComparisonOpInvalid(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.NegateToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	_, err := parser.ParseComparisonExp(0)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// if _, ok := tkn.(*token.AndToken); ok {
// 	return NewParseResult[Operator](&AndOp{}, position+1), nil
// } else if _, ok := tkn.(*token.OrToken); ok {
// 	return NewParseResult[Operator](&OrOp{}, position+1), nil
// } else if _, ok := tkn.(*token.NegateToken); ok {
// 	return NewParseResult[Operator](&NegateOp{}, position+1), nil

func TestLogicalExpAnd(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.AndToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseLogicalExp(0)

	if !parseResult.Equals(NewParseResult[Exp](NewOpExp(&IntLiteralExp{1}, &AndOp{}, &IntLiteralExp{1}), 3)) {

		t.Error("Parse result did not equal the expected result")
	}
}

func TestLogicalOpOr(t *testing.T) {
	tokens := []token.Token{&token.OrToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseLogicalOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&OrOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}
}

func TestLogicalExpOr(t *testing.T) {
	tokens := []token.Token{&token.IntLiteralToken{Number: 1}, &token.OrToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseLogicalExp(0)

	if !parseResult.Equals(NewParseResult[Exp](NewOpExp(&IntLiteralExp{1}, &OrOp{}, &IntLiteralExp{1}), 3)) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestParseStmtTestExp(t *testing.T) {
	tokens := []token.Token{&token.VariableToken{Name: "x"}, &token.EqualsToken{}, &token.IntLiteralToken{Number: 1}}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseStmt(0)

	if parseResult.Equals(NewParseResult[Stmt]) {
		t.Error("Parse result did not equal the expected result")
	}
}

func TestIfStmt(t *testing.T) {
	tokens := []token.Token{
		&token.IfToken{},
		&token.LeftParenToken{},
		&token.IntLiteralToken{1},
		&token.EqualsToken{},
		&token.IntLiteralToken{1},
		&token.RightParenToken{},
		&token.LeftCurlyToken{},
		&token.RightCurlyToken{},
	}

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseStmt(0)

	expected := NewParseResult[Stmt](NewIfStmt(
		NewOpExp(&IntLiteralExp{1}, &EqualsOp{}, &IntLiteralExp{1}),
		NewBlockStmt([]Stmt{})), 7)

	if !parseResult.Equals(expected) {
		t.Error("expected parse result does not equal actual")
	}
}

func TestIfElseStmt(t *testing.T) {
	tokens := tokenize("if(1==1){} else{}")

	parser := NewParser(tokens)
	parseResult, _ := parser.ParseStmt(0)

	expected := NewParseResult[Stmt](NewIfElseStmt(
		NewOpExp(&IntLiteralExp{1}, &EqualsOp{}, &IntLiteralExp{1}),
		NewBlockStmt([]Stmt{}),
		NewBlockStmt([]Stmt{}),
	), 10)

	if !parseResult.Equals(expected) {
		t.Error("expected parse result does not equal actual")
	}
}

func TestFuncDef(t *testing.T) {
	tokens := tokenize("fn int test(x int){}")
	parser := NewParser(tokens)
	expected := NewParseResult(NewFunctionDef(
		NewFunctionName("test"),
		[]*Vardec{NewVardec(NewVariable("x"), &IntType{})},
		NewBlockStmt([]Stmt{}),
		&IntType{},
	), len(tokens)-1)

	parseResult, err := parser.ParseFunctionDefinition(0)
	if parseResult == nil || err != nil {
		t.Fatalf("program did not parse")
	}
	if !parseResult.Equals(expected) {
		t.Fatalf("incorrect AST for ParseFunctionDefinition function")
	}

}
func TestProgram(t *testing.T) {
	tokens := tokenize(`
		fn int funcName(x int){}
		fn int anotherFunc(a int){}
		fn int thirdFunc(a int, b int){}`)

	parser := NewParser(tokens)
	if parseResult, err := parser.ParseProgram(); parseResult == nil || err != nil {
		t.Fatalf("program did not parse")
	}

}

// func TestFunctionDefinitionStmt(t *testing.T) {

// 	tokens := tokenize("fn")

// 	parser := NewParser(tokens)
// 	parseResult, _ := parser.parseFunctionDefinition(0)
// 	t.Logf("%#v", parseResult)

// 	expected := NewParseResult[Stmt](NewFunctionDef(*NewFunctionName("lol"),
// 		[]Vardec{{Name: "abc"}},
// 		NewBlockStmt([]Stmt{}),
// 		&IntType{},
// 	), 0)

// }
