package token

import "reflect"

type LesserToken struct {
}

func (lt *LesserToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(lt)
}

func (lt *LesserToken) String() string {
	return "<"
}
