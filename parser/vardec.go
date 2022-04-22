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

func (VD *Vardec) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherVardec, ok := reflect.Indirect(castOther).Interface().(Vardec)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(VD) && VD.Variable.Equals(otherVardec.Variable) && VD.Type.Equals(otherVardec.Type)
}
