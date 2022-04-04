package parser

import (
	"reflect"
	"strconv"
)

type NumberExp struct {
	Number int
}

func (ne *NumberExp) Equals(other interface{}) bool {
	if reflect.TypeOf(ne) == reflect.TypeOf(other) {
		castNumExp := reflect.ValueOf(other)
		castNumExp = reflect.Indirect(castNumExp)
		otherNumber := int(castNumExp.FieldByName("Number").Int())
		return otherNumber == ne.Number
	}
	return false

}

func (ne *NumberExp) String() string {
	return strconv.Itoa(ne.Number)
}

func (ne *NumberExp) exp() {}
