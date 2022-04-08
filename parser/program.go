package parser

type Program struct {
	Stmts Stmt
}

func NewProgram(stmts Stmt) *Program {
	return &Program{Stmts: stmts}
}
