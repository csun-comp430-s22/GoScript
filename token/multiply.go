package token

import "reflect"

type MultToken struct {
}

func (mt *MultToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mt)
}

func (mt *MultToken) String() string {
	return "*"
}
