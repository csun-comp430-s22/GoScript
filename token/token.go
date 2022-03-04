package token

type Token interface {
	Equals(other interface{}) bool
	String() string
}
