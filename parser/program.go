package parser

type Program struct {
	Stmt Stmt
	Node
}

func NewProgram(stmt Stmt) *Program {
	return &Program{Stmt: stmt}
}

func (p *Program) Equals(other any) bool {
	return true
}

func (p *Program) stmt() {}
