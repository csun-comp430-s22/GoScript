package token

import (
	"reflect"
	"strconv"
)

type FloatNumberToken struct {
	firstNumberHalf  float64
	secondNumberHalf float64
}

func (fnt *FloatNumberToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(fnt)
}

func (fnt *FloatNumberToken) String() string {
	return strconv.FormatFloat(fnt.firstNumberHalf, 'f', -1, 64) + "." + strconv.FormatFloat(fnt.secondNumberHalf, 'f', -1, 64)
}
