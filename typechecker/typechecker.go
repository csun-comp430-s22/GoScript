package typechecker

import (
	"fmt"

	"github.com/vSterlin/goscript/parser"
)

type TypeEnvironment map[parser.Variable]parser.Type

type Typechecker struct {
	functions []*parser.FunctionDef
}

func NewTypechecker(program parser.Program) *Typechecker {
	return &Typechecker{}
}

func (t *Typechecker) TypeOf(exp parser.Exp, typeEnv TypeEnvironment) (parser.Type, error) {

	if _, ok := exp.(*parser.IntLiteralExp); ok {
		return &parser.IntType{}, nil
	} else if varExp, ok := exp.(*parser.VariableExp); ok {
		variable := varExp.Variable
		varType := typeEnv[variable]
		// not in scope
		if varType == nil {
			panic(NewTypecheckerError("variable not in scope: " + variable.Name))
		} else {
			return varType, nil
		}
	} else if opExp, ok := exp.(*parser.OperatorExp); ok {
		return t.TypeOfOpExp(*opExp, typeEnv)
	} else if funcCallExp, ok := exp.(*parser.FunctionCallExp); ok {
		fmt.Printf("funcCallExp: %v\n", funcCallExp)
	}

	return nil, NewTypecheckerError("typechecker err")
}

func (t *Typechecker) TypeOfOpExp(exp parser.OperatorExp, typeEnv TypeEnvironment) (parser.Type, error) {

	leftType, _ := t.TypeOf(exp.Left, typeEnv)
	rightType, _ := t.TypeOf(exp.Right, typeEnv)

	if _, ok := exp.Op.(*parser.PlusOp); ok {
		_, leftOk := leftType.(*parser.IntType)
		_, rightOk := rightType.(*parser.IntType)
		if leftOk && rightOk {
			return &parser.IntType{}, nil
		} else {
			return nil, NewTypecheckerError("incorrect types for +")
		}

	} else if _, ok := exp.Op.(*parser.LessOp); ok {
		_, leftOk := leftType.(*parser.IntType)
		_, rightOk := rightType.(*parser.IntType)
		if leftOk && rightOk {
			return &parser.BoolType{}, nil
		} else {
			return nil, NewTypecheckerError("incorrect types for <")
		}
	} else if _, ok := exp.Op.(*parser.AndOp); ok {
		_, leftOk := leftType.(*parser.BoolType)
		_, rightOk := rightType.(*parser.BoolType)
		if leftOk && rightOk {
			return &parser.BoolType{}, nil
		} else {
			return nil, NewTypecheckerError("incorrect types for &&")
		}
	} else {
		return nil, NewTypecheckerError(fmt.Sprintf("unsupported operation: %#v", exp.Op))

	}

}

// func (t *Typechecker) TypeOfFuncCallExp(exp parser.FunctionCallExp, typEnv TypeEnvironment) (parser.Type, error) {

// }
