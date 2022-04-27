package token

import "reflect"

type HigherOrderToken struct {
}

func (hot *HigherOrderToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(hot)
}

func (hot *HigherOrderToken) String() string {
	return "higherorder"
}
