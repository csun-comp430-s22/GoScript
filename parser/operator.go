package parser

type Operator interface {
	Equals(other interface{}) bool
	Node
}
