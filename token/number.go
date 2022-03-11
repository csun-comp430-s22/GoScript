package token

import (
	"reflect"
	"strconv"
)

type NumberToken struct {
	Number int
}

func (nt *NumberToken) Equals(other interface{}) bool {

	castNumToken := reflect.ValueOf(other)
	castNumToken = reflect.Indirect(castNumToken)

	otherNumber := int(castNumToken.FieldByName("Number").Int())

	return reflect.TypeOf(*nt).Name() == castNumToken.Type().Name() && otherNumber == nt.Number
}

func (nt *NumberToken) String() string {
	return strconv.Itoa(nt.Number)
}
