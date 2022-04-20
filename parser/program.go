package parser

type Program struct {
	Stmt Stmt
	Node
}

func NewProgram(stmt Stmt) *Program {
	return &Program{Stmt: stmt}
}

func (p *Program) stmt() {}
