package token

import "reflect"

type ReturnToken struct {
}

func (rtt *ReturnToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(rtt)
}

func (rtt *ReturnToken) String() string {
	return "return"
}
