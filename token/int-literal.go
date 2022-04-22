package token

import (
	"reflect"
	"strconv"
)

type IntLiteralToken struct {
	Number int
}

func (nt *IntLiteralToken) Equals(other interface{}) bool {

	if reflect.TypeOf(nt) == reflect.TypeOf(other) {
		castNumToken := reflect.ValueOf(other)
		castNumToken = reflect.Indirect(castNumToken)
		otherNumber := int(castNumToken.FieldByName("Number").Int())
		return otherNumber == nt.Number
	}
	return false
}

func (nt *IntLiteralToken) String() string {
	return strconv.Itoa(nt.Number)
}
