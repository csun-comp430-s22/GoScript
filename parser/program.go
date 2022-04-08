package parser

type Program struct {
	Stmts Stmt
}

func NewProgramOp(stmts Stmt) *Program {
	return &Program{Stmts: stmts}
}
