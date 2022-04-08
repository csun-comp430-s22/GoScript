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

/* func TestExponentOpPower(t *testing.T) {

	tokens := []token.Token{&token.PowerToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&PowOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

} */

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

/* func TestModuloOpMod(t *testing.T) {

	tokens := []token.Token{&token.ModToken{}}
	parser := NewParser(tokens)
	parseResult, err := parser.ParseAdditiveOp(0)
	if err != nil {
		t.Error("Unexpected parser error")
	}

	if !parseResult.Equals(NewParseResult[Operator](&ModOp{}, 1)) {
		t.Error("Expected parse result did not equal actual")

	}

} */

func TestLessThanExp(t *testing.T) {
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
}
