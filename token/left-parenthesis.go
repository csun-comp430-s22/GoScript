package token

import "reflect"

type LeftParenToken struct {
}

func (lpt *LeftParenToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(lpt)
}

func (lpt *LeftParenToken) HashCode() int {
	return 2
}

func (lpt *LeftParenToken) String() string {
	return "("
}
