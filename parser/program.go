package parser

// type Program struct {
// 	Stmt Stmt
// 	Node
// }

// func NewProgram(stmt Stmt) *Program {
// 	return &Program{Stmt: stmt}
// }

type Program struct {
	FuncDefs []*FunctionDef
	Node
}

func NewProgram(funcDefs []*FunctionDef) *Program {
	return &Program{FuncDefs: funcDefs}
}

func (p *Program) stmt() {}
