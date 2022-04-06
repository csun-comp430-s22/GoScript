package parser

type Stmt interface {
	stmt()
	Node
}
