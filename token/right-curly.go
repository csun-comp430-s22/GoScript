package token

import "reflect"

type RightCurlyToken struct {
}

func (rct *RightCurlyToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(rct)
}

func (rct *RightCurlyToken) HashCode() int {
	return 6
}

func (rct *RightCurlyToken) String() string {
	return "{"
}
