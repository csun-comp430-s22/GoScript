package token

import "reflect"

type FalseToken struct {
}

func (ft *FalseToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(ft)
}

func (ft *FalseToken) String() string {
	return "false"
}
