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

func TestVariableIsInteger(t *testing.T) {
	typechecker := NewTypechecker(*parser.NewProgram([]*parser.FunctionDef{}))

	intVar := parser.NewVardec(parser.NewVariable("test"), &parser.IntType{})

	typeEnv := TypeEnvironment{*intVar.Variable: &parser.IntType{}}

	actualType, _ := typechecker.TypeOf(parser.NewVariableExp(*intVar.Variable), typeEnv)
	expectedType := &parser.IntType{}

	if !actualType.Equals(expectedType) {
		t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
	}

}

func TestVariableOutOfScope(t *testing.T) {
	typechecker := NewTypechecker(*parser.NewProgram([]*parser.FunctionDef{}))

	intVar := parser.NewVardec(parser.NewVariable("test"), &parser.IntType{})

	typeEnv := TypeEnvironment{}

	_, err := typechecker.TypeOf(parser.NewVariableExp(*intVar.Variable), typeEnv)

	if err == nil {
		t.Errorf("expected error to notify about variable out of scope")
	}

}

func TestPlusOpExp(t *testing.T) {

	plusOpExp := parser.NewOpExp(&parser.IntLiteralExp{Number: 1}, &parser.PlusOp{}, &parser.IntLiteralExp{Number: 1})

	actualType, _ := typeOf(plusOpExp)
	expectedType := &parser.IntType{}

	if !actualType.Equals(expectedType) {
		t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
	}
}

func TestComparisonOpExp(t *testing.T) {

	compOps := []parser.Type{
		&parser.LessOp{},
		&parser.LessEqualOp{},
		&parser.GreaterOp{},
		&parser.GreaterEqualOp{},
		&parser.EqualsOp{},
		&parser.NotEqualsOp{},
	}

	for _, compOp := range compOps {
		compOpExp := parser.NewOpExp(&parser.IntLiteralExp{Number: 1}, compOp, &parser.IntLiteralExp{Number: 1})

		actualType, _ := typeOf(compOpExp)
		expectedType := &parser.BoolType{}

		if !actualType.Equals(expectedType) {
			t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
		}
	}
}