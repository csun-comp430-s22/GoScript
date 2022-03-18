package token

import (
	"reflect"
	"strconv"
)

type NumberToken struct {
	Number int
}

func (nt *NumberToken) Equals(other interface{}) bool {

	if reflect.TypeOf(nt) == reflect.TypeOf(other) {
		castNumToken := reflect.ValueOf(other)
		castNumToken = reflect.Indirect(castNumToken)
		otherNumber := int(castNumToken.FieldByName("Number").Int())
		return otherNumber == nt.Number
	}
	return false
}

func (nt *NumberToken) String() string {
	return strconv.Itoa(nt.Number)
}
