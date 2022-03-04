package token

import "reflect"

type QuoteStringToken struct {
	name string
}

func (qst *QuoteStringToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(qst)
}

func (qst *QuoteStringToken) String() string {
	return "\"" + qst.name + "\""
}
