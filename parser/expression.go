package parser

// maybe add some kind of Exp() method
// to differentiate between other empty interfaces
type Exp interface {
	exp()
	Node
}
