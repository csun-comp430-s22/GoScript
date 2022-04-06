package parser

type Node interface {
	Equals(other interface{}) bool
}
