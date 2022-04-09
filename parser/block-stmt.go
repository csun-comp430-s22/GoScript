package parser

type BlockStmt struct {
	Stmts []Stmt
}

func NewBlockStmt(stmts []Stmt) *BlockStmt {
	return &BlockStmt{Stmts: stmts}
}

// TODO
func (bs *BlockStmt) Equals(other any) bool {
	return true
}

func (bs *BlockStmt) stmt() {}
