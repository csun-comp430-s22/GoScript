package parser

import (
	"reflect"
)

type Vardec struct {
	Variable *Variable
	Type     Type
}

func NewVardec(v *Variable, t Type) *Vardec {
	return &Vardec{Variable: v, Type: t}
}

func (vd *Vardec) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherVardec, ok := reflect.Indirect(castOther).Interface().(Vardec)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(vd) && vd.Variable.Equals(otherVardec.Variable) && vd.Type.Equals(otherVardec.Type)
}
