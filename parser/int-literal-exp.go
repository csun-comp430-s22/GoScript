package parser

import (
	"reflect"
)

type IntLiteralExp struct {
	Number int
}

func (ile *IntLiteralExp) Equals(other interface{}) bool {
	if reflect.TypeOf(ile) == reflect.TypeOf(other) {
		castNumExp := reflect.ValueOf(other)
		castNumExp = reflect.Indirect(castNumExp)
		otherNumber := int(castNumExp.FieldByName("Number").Int())
		return otherNumber == ile.Number
	}
	return false

}

func (ile *IntLiteralExp) exp() {}
