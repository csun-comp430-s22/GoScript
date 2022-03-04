package token

import "reflect"

type QuoteToken struct {
}

func (qtt *QuoteToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(qtt)
}

func (qtt *QuoteToken) String() string {
	return "\""
}
