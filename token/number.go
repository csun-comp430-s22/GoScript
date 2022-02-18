package token

import (
	"reflect"
	"strconv"
)

type NumberToken struct {
	number int
}

func (nt *NumberToken) HashCode() int {
	return nt.number
}

func (nt *NumberToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(nt)
}

func (nt *NumberToken) String() string {
	return strconv.Itoa(nt.number)
}
