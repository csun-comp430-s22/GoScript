package token

type Token interface {
	Equals(other interface{}) bool
	HashCode() int // probably don't need this in go
	String() string
}
