package typechecker

import (
	"testing"

	"github.com/vSterlin/goscript/parser"
)

func typeOf(exp parser.Exp) (parser.Type, error) {
	emptyTypeChecker := NewTypechecker(*parser.NewProgram([]*parser.FunctionDef{}))

	return emptyTypeChecker.TypeOf(exp, TypeEnvironment{})
}

func TestTypecheckInt(t *testing.T) {

	actualType, _ := typeOf(&parser.IntLiteralExp{Number: 0})
	expectedType := &parser.IntType{}

	if !actualType.Equals(expectedType) {
		t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
	}
}

func TestTypecheckBool(t *testing.T) {

	actualType, _ := typeOf(&parser.BoolLiteralExp{Value: true})
	expectedType := &parser.BoolType{}

	if !actualType.Equals(expectedType) {
		t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
	}
}
