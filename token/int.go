package token

import "reflect"

type IntToken struct {
}

func (int *IntToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(int)
}

func (int *IntToken) String() string {
	return "int"
}
