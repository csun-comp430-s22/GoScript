package parser

// maybe add some kind of Exp() method
// to differentiate between other empty interfaces
type Exp interface {
	Equals(other interface{}) bool
	String() string
}
