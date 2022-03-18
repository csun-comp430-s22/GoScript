package parser

type Operator interface {
	Equals(other interface{}) bool
	String() string
}
