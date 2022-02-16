package token

import "reflect"

type IfToken struct {
}

func (it *IfToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(it)
}

func (it *IfToken) HashCode() int {
	return 4
}

func (it *IfToken) String() string {
	return "if"
}
