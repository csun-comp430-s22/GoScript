package parser

import (
	"reflect"
)

type BoolLiteralExp struct {
	Value bool
}

func (ble *BoolLiteralExp) Equals(other interface{}) bool {
	if reflect.TypeOf(ble) == reflect.TypeOf(other) {
		castBoolExp := reflect.ValueOf(other)
		castBoolExp = reflect.Indirect(castBoolExp)
		otherValue := bool(castBoolExp.FieldByName("Value").Bool())
		return otherValue == ble.Value
	}
	return false

}

func (ble *BoolLiteralExp) exp() {}
