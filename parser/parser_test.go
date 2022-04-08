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
