package token

import "reflect"

type ModToken struct {
}

func (mdt *ModToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mdt)
}

func (mdt *ModToken) String() string {
	return "%"
}
