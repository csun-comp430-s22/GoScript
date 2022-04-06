package parser

type Stmt interface {
	Equals(other interface{}) bool
	Stmt()
	Node
}
