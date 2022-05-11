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

	intVar := parser.NewVardec(parser.NewVariable("test", false), &parser.IntType{})

	typeEnv := TypeEnvironment{*intVar.Variable: &parser.IntType{}}

	actualType, _ := typechecker.TypeOf(parser.NewVariableExp(*intVar.Variable), typeEnv)
	expectedType := &parser.IntType{}

	if !actualType.Equals(expectedType) {
		t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
	}

}

func TestVariableOutOfScope(t *testing.T) {
	typechecker := NewTypechecker(*parser.NewProgram([]*parser.FunctionDef{}))

	intVar := parser.NewVardec(parser.NewVariable("test", false), &parser.IntType{})

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

func TestLogicalOpExp(t *testing.T) {

	logicalOps := []parser.Type{
		&parser.AndOp{},
		&parser.OrOp{},
		&parser.NegateOp{},
	}

	for _, logicalOp := range logicalOps {
		logicalOpExp := parser.NewOpExp(&parser.BoolLiteralExp{Value: true}, logicalOp, &parser.BoolLiteralExp{Value: false})

		actualType, _ := typeOf(logicalOpExp)
		expectedType := &parser.BoolType{}

		if !actualType.Equals(expectedType) {
			t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
		}
	}
}

func TestFunctionCallExpType(t *testing.T) {

	funcName := parser.NewFunctionName("test")
	funcDef := parser.NewFunctionDef(funcName, []*parser.Vardec{}, parser.NewBlockStmt([]parser.Stmt{}), &parser.IntType{})

	funcCallExp := parser.NewFunctionCallExp(funcName, []parser.Exp{})

	typechecker := NewTypechecker(*parser.NewProgram([]*parser.FunctionDef{funcDef}))

	actualType, _ := typechecker.TypeOf(funcCallExp, TypeEnvironment{})
	expectedType := &parser.IntType{}

	if !actualType.Equals(expectedType) {
		t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
	}

}

func TestFunctionCallExpWithArgsType(t *testing.T) {

	funcName := parser.NewFunctionName("test")
	funcArgs := []*parser.Vardec{parser.NewVardec(parser.NewVariable("a", false), &parser.IntType{})}
	funcDef := parser.NewFunctionDef(funcName, funcArgs, parser.NewBlockStmt([]parser.Stmt{}), &parser.IntType{})

	funcCallExp := parser.NewFunctionCallExp(funcName, []parser.Exp{&parser.IntLiteralExp{Number: 1}})

	typechecker := NewTypechecker(*parser.NewProgram([]*parser.FunctionDef{funcDef}))

	actualType, _ := typechecker.TypeOf(funcCallExp, TypeEnvironment{})
	expectedType := &parser.IntType{}

	if !actualType.Equals(expectedType) {
		t.Errorf("types mismatch: %#v type does not equal %#v type", actualType, expectedType)
	}

}

func TestFunctionCallExpWithArgsMismatch(t *testing.T) {

	funcName := parser.NewFunctionName("test")
	funcArgs := []*parser.Vardec{parser.NewVardec(parser.NewVariable("a", false), &parser.IntType{})}
	funcDef := parser.NewFunctionDef(funcName, funcArgs, parser.NewBlockStmt([]parser.Stmt{}), &parser.BoolType{})

	funcCallExp := parser.NewFunctionCallExp(funcName, []parser.Exp{&parser.IntLiteralExp{Number: 1}})

	typechecker := NewTypechecker(*parser.NewProgram([]*parser.FunctionDef{funcDef}))

	actualType, _ := typechecker.TypeOf(funcCallExp, TypeEnvironment{})
	expectedType := &parser.IntType{}

	if actualType.Equals(expectedType) {
		t.Errorf("types were found to be equal: %#v type should not match %#v type", actualType, expectedType)
	}

}

type unsupportedOp struct{}

func (uo *unsupportedOp) Equals(other any) bool {
	return false
}

func TestUnsupportedOperation(t *testing.T) {

	logicalOpExp := parser.NewOpExp(&parser.BoolLiteralExp{Value: true}, &unsupportedOp{}, &parser.BoolLiteralExp{Value: false})

	_, err := typeOf(logicalOpExp)

	if err == nil {
		t.Errorf("expected error to notify about unusupported operator")
	}

}
