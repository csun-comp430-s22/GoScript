package token

import "reflect"

type NotEqualsToken struct {
}

func (net *NotEqualsToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(net)
}

func (net *NotEqualsToken) String() string {
	return "!="
}
