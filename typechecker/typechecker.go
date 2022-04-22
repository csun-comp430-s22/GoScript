package typechecker

import "github.com/vSterlin/goscript/parser"

type Typechecker struct {
}

func (t *Typechecker) TypeOf(exp parser.Exp, typeEnv map[parser.Variable]parser.Type) (parser.Type, error) {

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
	}

	return nil, NewTypecheckerError("typechecker err")
}

func (t *Typechecker) TypeOfOpExp(exp parser.OperatorExp, typeEnv map[parser.Variable]parser.Type) (parser.Type, error) {

	leftType, _ := t.TypeOf(exp.Left, typeEnv)
	rightType, _ := t.TypeOf(exp.Right, typeEnv)

	if _, ok := exp.Op.(*parser.PlusOp); ok {
		_, leftOk := leftType.(*parser.IntType)
		_, rightOk := rightType.(*parser.IntType)
		if leftOk && rightOk {
			return &parser.IntType{}, nil
		}
	}

	// for now
	return nil, nil

}
