package parser

type Type interface {
	Equals(other interface{}) bool
}
