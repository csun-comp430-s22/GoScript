package token

import "reflect"

type ModToken struct {
}

func (mt *ModToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mt)
}

func (mt *ModToken) String() string {
	return "%"
}
