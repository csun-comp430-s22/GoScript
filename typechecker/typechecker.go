package typechecker

import (
	"fmt"

	"github.com/vSterlin/goscript/parser"
)

type TypeEnvironment map[parser.Variable]parser.Type

type Typechecker struct {
	Functions []*parser.FunctionDef
}

func NewTypechecker(program parser.Program) *Typechecker {
	return &Typechecker{Functions: program.FuncDefs}
}

// // rewritten below with nicer syntax
// func (t *Typechecker) TypeOf(exp parser.Exp, typeEnv TypeEnvironment) (parser.Type, error) {

// 	if _, ok := exp.(*parser.IntLiteralExp); ok {
// 		return &parser.IntType{}, nil
// 	} else if _, ok := exp.(*parser.BoolLiteralExp); ok {
// 		return &parser.BoolType{}, nil
// 	} else if varExp, ok := exp.(*parser.VariableExp); ok {
// 		variable := varExp.Variable
// 		varType := typeEnv[variable]
// 		// not in scope
// 		if varType == nil {
// 			panic(NewTypecheckerError("variable not in scope: " + variable.Name))
// 		} else {
// 			return varType, nil
// 		}
// 	} else if opExp, ok := exp.(*parser.OperatorExp); ok {
// 		return t.TypeOfOpExp(*opExp, typeEnv)
// 	} else if funcCallExp, ok := exp.(*parser.FunctionCallExp); ok {
// 		fmt.Printf("funcCallExp: %v\n", funcCallExp)
// 	}

// 	return nil, NewTypecheckerError("typechecker err")
// }

func (t *Typechecker) TypeOf(exp parser.Exp, typeEnv TypeEnvironment) (parser.Type, error) {

	switch castExp := exp.(type) {
	case (*parser.IntLiteralExp):
		return &parser.IntType{}, nil

	case (*parser.BoolLiteralExp):
		return &parser.BoolType{}, nil

	case (*parser.VariableExp):
		variable := castExp.Variable
		varType := typeEnv[variable]
		// not in scope
		if varType == nil {
			return nil, NewTypecheckerError("variable not in scope: " + variable.Name)
		} else {
			return varType, nil
		}

	case (*parser.OperatorExp):
		return t.TypeOfOpExp(*castExp, typeEnv)

	case (*parser.FunctionCallExp):
		return t.TypeOfFuncCallExp(*castExp, typeEnv)
	}
	return nil, NewTypecheckerError("typechecker err")
}

func (t *Typechecker) TypeOfOpExp(exp parser.OperatorExp, typeEnv TypeEnvironment) (parser.Type, error) {

	leftType, _ := t.TypeOf(exp.Left, typeEnv)
	rightType, _ := t.TypeOf(exp.Right, typeEnv)

	// additive
	if _, ok := exp.Op.(*parser.PlusOp); ok {
		_, leftOk := leftType.(*parser.IntType)
		_, rightOk := rightType.(*parser.IntType)
		if leftOk && rightOk {
			return &parser.IntType{}, nil
		} else {
			return nil, NewTypecheckerError("incorrect types for +")
		}

		// comparison
	} else if comparisonOpType := tryTypeOfComparisonExp(exp); comparisonOpType != nil {
		_, leftOk := leftType.(*parser.IntType)
		_, rightOk := rightType.(*parser.IntType)
		if leftOk && rightOk {
			return &parser.BoolType{}, nil
		} else {
			return nil, NewTypecheckerError(fmt.Sprintf("incorrect types for comparison operator: %#v", comparisonOpType))
		}
	} else if logicalOpType := tryTypeofLogicalExp(exp); logicalOpType != nil {
		_, leftOk := leftType.(*parser.BoolType)
		_, rightOk := rightType.(*parser.BoolType)
		if leftOk && rightOk {
			return &parser.BoolType{}, nil
		} else {
			return nil, NewTypecheckerError(fmt.Sprintf("incorrect types for logical operator: %#v", logicalOpType))
		}
	} else {
		return nil, NewTypecheckerError(fmt.Sprintf("unsupported operation: %#v", exp.Op))

	}

}

func tryTypeOfComparisonExp(exp parser.OperatorExp) parser.Type {

	compOpTypes := []parser.Type{
		&parser.LessOp{},
		&parser.LessEqualOp{},
		&parser.GreaterOp{},
		&parser.GreaterEqualOp{},
		&parser.EqualsOp{},
		&parser.NotEqualsOp{},
	}

	for _, compOpType := range compOpTypes {
		if exp.Op.Equals(compOpType) {
			return exp.Op
		}
	}

	return nil

}

func tryTypeofLogicalExp(exp parser.OperatorExp) parser.Type {

	logicalOpTypes := []parser.Type{
		&parser.AndOp{},
		&parser.OrOp{},
		&parser.NegateOp{},
	}

	for _, logicalOpType := range logicalOpTypes {
		if exp.Op.Equals(logicalOpType) {
			return exp.Op
		}
	}

	return nil
}

func (t *Typechecker) GetFunctionByName(funcName *parser.FunctionName) (*parser.FunctionDef, error) {
	for _, fDef := range t.Functions {
		if fDef.Name.Equals(funcName) {
			return fDef, nil
		}
	}

	return nil, NewTypecheckerError("No function with name: " + funcName.Name)
}

func (t *Typechecker) TypeOfFuncCallExp(exp parser.FunctionCallExp, typeEnv TypeEnvironment) (parser.Type, error) {

	fDef, err := t.GetFunctionByName(exp.FunctionName)
	if err != nil {
		return nil, err
	}

	if len(exp.Params) != len(fDef.Args) {
		return nil, NewTypecheckerError("Wrong number of arguments for function: " + fDef.Name.Name)
	}

	for i, param := range exp.Params {
		receivedArgType, err := t.TypeOf(param, typeEnv)
		if err != nil {
			return nil, err
		}
		expectedArgType := fDef.Args[i].Type
		if !receivedArgType.Equals(expectedArgType) {
			return nil, NewTypecheckerError("Type mismatch on function call argument")
		}
	}

	return fDef.ReturnType, nil

}
