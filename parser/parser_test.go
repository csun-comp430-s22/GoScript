package parser

import (
	"fmt"

	"testing"

	"github.com/vSterlin/goscript/token"
)

func TestPlusOpExp(t *testing.T) {
	first := NewOpExp(&NumberExp{1}, &PlusOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &PlusOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &PlusOp{}, &NumberExp{2})

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

	tokens := []token.Token{&token.PlusToken{}}
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
	first := NewOpExp(&NumberExp{1}, &MinusOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &MinusOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &MinusOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &DivideOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &DivideOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &DivideOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &MultiplyOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &MultiplyOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &MultiplyOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &PowOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &PowOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &PowOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &ModOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &ModOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &ModOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &AndOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &AndOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &AndOp{}, &NumberExp{2})

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
	tokens := []token.Token{&token.NumberToken{1}, &token.LesserToken{}, &token.NumberToken{2}}
	parser := NewParser(tokens)

	parseResult, err := parser.ParseComparisonExp(0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	expected := NewParseResult(NewOpExp(&NumberExp{1}, &LessOp{}, &NumberExp{2}), 3)

	fmt.Printf("parseResult: %v\n", parseResult)
	fmt.Printf("expected: %v\n", expected)

	if !parseResult.Equals(expected) {
		t.Error("Expected parse result did not equal actual")

	}
} */

func TestPipeOperatorOpExp(t *testing.T) {
	first := NewOpExp(&NumberExp{1}, &PipeOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &PipeOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &PipeOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &EqualsOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &EqualsOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &EqualsOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &NotEqualsOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &NotEqualsOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &NotEqualsOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &OrOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &OrOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &OrOp{}, &NumberExp{2})

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
	parseResult, err := parser.ParseLogicalOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&OrOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

}

func TestLessThanOp(t *testing.T) {
	first := NewOpExp(&NumberExp{1}, &LessOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &LessOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &LessOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &AndOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &AndOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &AndOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &GreaterOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &GreaterOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &GreaterOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &LessEqualOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &LessEqualOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &LessEqualOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &GreaterEqualOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &GreaterEqualOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &GreaterEqualOp{}, &NumberExp{2})

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
	first := NewOpExp(&NumberExp{1}, &NegateOp{}, &NumberExp{1})
	second := NewOpExp(&NumberExp{1}, &NegateOp{}, &NumberExp{1})
	third := NewOpExp(&NumberExp{1}, &NegateOp{}, &NumberExp{2})

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
