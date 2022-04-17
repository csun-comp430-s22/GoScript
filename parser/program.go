package parser

type Program struct {
	Stmts Stmt
	Node
}

func NewProgram(stmts Stmt) *Program {
	return &Program{Stmts: stmts}
}

func (p *Program) stmt() {}
