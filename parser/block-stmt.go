package parser

type BlockStmt struct {
	Stmts []Stmt
}

func NewBlockStmtOp(stmts []Stmt) *BlockStmt {
	return &BlockStmt{Stmts: stmts}
}
