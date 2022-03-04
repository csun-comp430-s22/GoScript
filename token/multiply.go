package token

import "reflect"

type MultToken struct {
}

func (mlt *MultToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mlt)
}

func (mlt *MultToken) String() string {
	return "*"
}
