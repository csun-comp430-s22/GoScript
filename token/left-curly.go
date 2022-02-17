package token

import "reflect"

type LeftCurlyToken struct {
}

func (lct *LeftCurlyToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(lct)
}

func (lct *LeftCurlyToken) HashCode() int {
	return 5
}

func (lct *LeftCurlyToken) String() string {
	return "{"
}
