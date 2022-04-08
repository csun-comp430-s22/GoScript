package parser

import (
	"fmt"
	"testing"
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
